# StickyHeader Implementation Review

## Overview
This document compares the Go-Compose [StickyHeader](file:///home/jaco/SecondBrain/1-Projects/GoCompose/development/go-compose/compose/foundation/lazy/lazy_list_scope.go#10-11) implementation in [lazy_list.go](file:///home/jaco/SecondBrain/1-Projects/GoCompose/development/go-compose/compose/foundation/lazy/lazy_list.go) with the reference Kotlin implementation in AndroidX [LazyLayoutStickyItems.kt](file:///home/jaco/SecondBrain/1-Projects/LearnIMGUI/clones/androidx/compose/foundation/foundation/src/commonMain/kotlin/androidx/compose/foundation/lazy/layout/LazyLayoutStickyItems.kt) and [LazyListMeasure.kt](file:///home/jaco/SecondBrain/1-Projects/LearnIMGUI/clones/androidx/compose/foundation/foundation/src/commonMain/kotlin/androidx/compose/foundation/lazy/LazyListMeasure.kt).

## Summary
The Go implementation of [StickyHeader](file:///home/jaco/SecondBrain/1-Projects/GoCompose/development/go-compose/compose/foundation/lazy/lazy_list_scope.go#10-11) successfully ports the core logic of "stick to top" and "push up" animations. However, there are some differences due to the underlying architecture (Gio's `layout.List` vs Jetpack Compose's custom `LazyLayout`) and some missing features (content padding support).

## Detailed Comparison

### 1. Sticky Header Logic
*   **Kotlin (`StickToTopPlacement`)**:
    *   Finds the active sticky header index (last header `<= firstVisibleItemIndex`).
    *   Places it at `-beforeContentPadding` (stick to top edge).
    *   Checks for the *next* visible sticky header to calculate collision.
    *   Adjusts offset: `minOf(stick_pos, next_header_pos - current_header_size)`.
*   **Go ([lazyListWidgetConstructor](file:///home/jaco/SecondBrain/1-Projects/GoCompose/development/go-compose/compose/foundation/lazy/lazy_list.go#85-243))**:
    *   Finds the active sticky header index (last header `<= state.List.List.Position.First`).
    *   Places it at `0` (stick to top edge).
    *   Iterates `itemSizes` (collected during current frame layout) to find the position of the *next* sticky header.
    *   Adjusts offset: `if pos < headerSize { headerOffset = pos - headerSize }`.
*   **Verdict**: **Functionally Equivalent**. The Go logic correctly implements the "push up" behavior using available layout information.

### 2. Implementation Differences & Limitations

#### A. Content Padding
*   **Kotlin**: Explicitly handles `beforeContentPadding` to ensure headers stick to the edge of the content area, not necessarily the absolute top of the container (if padding exists).
*   **Go**: Ignores any potential content padding (currently [StickyHeader](file:///home/jaco/SecondBrain/1-Projects/GoCompose/development/go-compose/compose/foundation/lazy/lazy_list_scope.go#10-11) sticks to `0`).
    *   *Impact*: If [LazyList](file:///home/jaco/SecondBrain/1-Projects/GoCompose/development/go-compose/compose/foundation/lazy/lazy_list_scope.go#7-12) supports content padding in the future, the header might look incorrect (covering padding). For now, as long as `contentPadding` isn't fully supported/used, this is acceptable.

#### B. Double Drawing (Efficiency/Alpha Blending)
*   **Kotlin**: Removes the sticky item from the normal `positionedItems` list so it is drawn *only* as a sticky item.
*   **Go**: The underlying `state.List.List.Layout` draws all visible items, including the one that is currently "sticky" (if it's within the viewport). The [StickyHeader](file:///home/jaco/SecondBrain/1-Projects/GoCompose/development/go-compose/compose/foundation/lazy/lazy_list_scope.go#10-11) logic then draws the header *again* on top.
    *   *Impact*:
        *   **Performance**: Slight overhead of drawing the same widget twice when it is naturally visible at the top.
        *   **Visuals**: If the header has transparency/alpha, it will appear darker/different because it is blended twice (once by list, once by sticky overlay).
    *   *Recommendation*: In the future, modify the list layout callback to suppress drawing (but keep measuring) the item if it identifies strictly as the active sticky header.

#### C. Next Header Visibility
*   **Difference**: Go relies on `itemSizes` map populated during the list layout. The "push up" logic only works if the *next* sticky header is effectively traversed/laid out by Gio.
*   **Analysis**: Since generic `layout.List` only processes visible items, if the next header is not visible, it won't be in `itemSizes`. The logic correctly handles this (`found = false`), resulting in no push-up. This is correct behavior (no collision if not visible).

#### D. Eager Composition
*   **Observation**: [LazyColumn](file:///home/jaco/SecondBrain/1-Projects/GoCompose/development/go-compose/compose/foundation/lazy/lazy_list.go#24-29) in [lazy_list.go](file:///home/jaco/SecondBrain/1-Projects/GoCompose/development/go-compose/compose/foundation/lazy/lazy_list.go) emits all items in `scope.items` as children to the layout node.
    ```go
    // Emit all items as children
    for _, item := range scope.items {
        c.WithComposable(item.Content)
    }
    ```
    *   *Impact*: This is a known limitation of the current Go port ("Eager Composition"). Unlike Kotlin's true lazy composition, this constructs the entire component tree description (though Gio might only layout visible ones). [StickyHeader](file:///home/jaco/SecondBrain/1-Projects/GoCompose/development/go-compose/compose/foundation/lazy/lazy_list_scope.go#10-11) performance is bound by this general architectural bottleneck.

## Conclusion
The [StickyHeader](file:///home/jaco/SecondBrain/1-Projects/GoCompose/development/go-compose/compose/foundation/lazy/lazy_list_scope.go#10-11) port is **correctly implemented** for the current stage of `go-compose`. It faithfully reproduces the visual behavior of the Kotlin reference.

**Action Items (Future):**
1.  **Fix Double Drawing**: Prevent the normal list from drawing the active sticky header to avoid alpha-blending artifacts.
2.  **Content Padding**: When `ContentPadding` support is added to [LazyList](file:///home/jaco/SecondBrain/1-Projects/GoCompose/development/go-compose/compose/foundation/lazy/lazy_list_scope.go#7-12), ensure [StickyHeader](file:///home/jaco/SecondBrain/1-Projects/GoCompose/development/go-compose/compose/foundation/lazy/lazy_list_scope.go#10-11) respects `beforeContentPadding`.