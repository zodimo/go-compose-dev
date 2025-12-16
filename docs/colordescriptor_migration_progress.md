# ColorDescriptor Migration Progress

This document tracks the migration of Material3 components to use the `ColorDescriptor` pattern with theme-aware defaults.

## Migration Status

Last Updated: 2025-12-16 (All Native Components Complete)

### ✅ Migrated to ColorDescriptor (19 components)

Components that use `theme.ColorDescriptor`:

| Component | Date Completed | Notes |
|-----------|----------------|-------|
| **AnimatedBackground** | 2025-12-16 | Modifier in `modifiers/animation`. Uses `ColorDescriptor`. |
| **surface** | - | Foundation component, uses ColorDescriptor for Color, ContentColor, BorderColor |
| **appbar** | 2025-12-16 | Uses ColorDescriptor for TopAppBarColors (5 fields: ContainerColor, ScrolledContainerColor, NavigationIconContentColor, TitleContentColor, ActionIconContentColor) |
| **bottomappbar** | 2025-12-16 | Uses ColorDescriptor for ContainerColor, ContentColor |
| **bottomsheet** | 2025-12-16 | Uses ColorDescriptor for ContainerColor, ScrimColor. Default: SurfaceContainerLow, Scrim. |
| **tab** | 2025-12-16 | Uses ColorDescriptor for TabRow ContainerColor/ContentColor and Tab SelectedContentColor/UnselectedContentColor |
| **scaffold** | 2025-12-16 | Uses ColorDescriptor for ContainerColor, ContentColor - core layout container |
| **floatingactionbutton** | 2025-12-16 | Uses ColorDescriptor for ContainerColor, ContentColor (defaults to PrimaryContainer/OnPrimaryContainer) |
| **chip** | 2025-12-16 | Uses ColorDescriptor for Color (Surface) and BorderColor (OutlineVariant) |
| **navigationbar** | 2025-12-16 | Uses ColorDescriptor for ContainerColor/ContentColor and Item IndicatorColor. Defaults: SurfaceContainer, OnVariant, SecondaryContainer |
| **segmentedbutton** | 2025-12-16 | Uses ColorDescriptor for Selected/Unselected Color/ContentColor and BorderColor. Defaults: SecondaryContainer, Surface, OnSecondaryContainer, OnSurface, Outline |
| **badge** | 2025-12-16 | Uses ColorDescriptor for ContainerColor, ContentColor. Defaults: Error, OnError |
| **slider** | 2025-12-16 | Uses ColorDescriptor for Thumb/Track/Tick colors. Defaults: Primary, OnPrimary, SurfaceContainerHighest, OnSurfaceVariant |
| **divider** | 2025-12-16 | Uses ColorDescriptor for Color (OutlineVariant). Native implementation. |
| **Foundation Text** | 2025-12-16 | Uses ColorDescriptor for Color, SelectionColor. Defaults: SurfaceRoles.OnSurface, PrimaryRoles.Primary with opacity |
| **Overlay** | 2025-12-16 | Uses ColorDescriptor for ScrimColor. Default: ScrimRoles.Scrim |
| **RadioButton** | 2025-12-16 | Uses ColorDescriptor for SelectedColor, UnselectedColor, DisabledColor. Defaults: PrimaryRoles.Primary, SurfaceRoles.OnVariant, SurfaceRoles.OnSurface with opacity |
| **Shadow** | 2025-12-16 | Modifier uses ColorDescriptor for AmbientColor, SpotColor. Default: ScrimRoles.Shadow |
| **NavigationDrawerItem** | 2025-12-16 | Uses ColorDescriptor for containerColor (selected: SecondaryRoles.Container) |
| **NavigationRail** | 2025-12-16 | Uses ColorDescriptor for ContainerColor, ContentColor. Defaults: SurfaceRoles.Surface, SurfaceRoles.OnSurface |

### ✅ All Native Components Migrated

All native go-compose components now use `theme.ColorDescriptor` with appropriate theme role defaults. No pending migrations remaining.


### 🔗 External Widget Components (12 confirmed)

Components that wrap gio-mw widgets and don't expose color options:

| Component | Widget/Package | Notes |
|-----------|----------------|-------|
| **button** | `gio-mw/widget/button` | No color options exposed |
| **card** | `gio-mw/widget/card` | No color options exposed |
| **checkbox** | `gio-mw/widget/checkbox` | No color options exposed |
| **dialog** | `gio-mw/widget/dialog` | No color options exposed |
| **iconbutton** | `gio-mw/widget/button` | No color options exposed |
| **menu** | `gio-mw/token` | Uses gio-mw tokens |
| **progress** | `gio-mw/widget/indicator` | Wrapper around gio-mw progress indicator |
| **snackbar** | `gio-mw/widget/snackbar` | Wrapper around gio-mw snackbar widget |
| **switch** | `gio-mw/widget/toggle` | No color options exposed |
| **text** | `gio-mw/token`, `gio-mw/wdk` | Uses gio-mw typography helpers |
| **textfield** | `gio-mw/widget/input` | Wrapper around gio-mw input widget |
| **tooltip** | `gio-mw/widget/tooltip` | Wrapper around gio-mw tooltip widget |

> [!NOTE]
> These components use the gio-mw widget library's Material theme system. They automatically respond to theme changes but do not expose color customization options.

> [!NOTE]
> No special cases remaining - all native components now categorized.

## Migration Workflow

For step-by-step migration instructions, see:
- [Migration Workflow](.agent/workflows/migrate_component_colordescriptor.md)

## Quick Reference

### Common Theme Role Mappings
- **Container backgrounds** → `SurfaceRoles.Surface*` (Surface, SurfaceContainer, SurfaceContainerHigh, etc.)
- **Content/text** → `ContentRoles.OnSurface`, `OnPrimary`, `OnSecondary`, etc.
- **Borders/outlines** → `OutlineRoles.Outline`, `OutlineVariant`
- **Primary actions** → `PrimaryRoles.Primary`, `PrimaryContainer`
- **State layers** → Use `.SetOpacity()` on base colors

### Migration Checklist
- [ ] Analyze color usage
- [ ] Update Options structure to use ColorDescriptor
- [ ] Set theme-aware defaults using color roles
- [ ] Update option setters to accept ColorDescriptor
- [ ] Remove SpecificColor() wrappers (keep only for non-theme colors)
- [ ] Update internal color resolution if needed
- [ ] Update tests/demos
- [ ] Verify build and visual appearance
- [ ] Update documentation

## Architecture Notes

### Complete Audit Summary (2025-12-16)

**Total Material3 Components: 28**

**ColorDescriptor Migration Progress: 13 of 19 native components/modifiers (68%)**

| Category | Count | Components |
|----------|-------|------------|
| ✅ **Migrated** | 13 | surface, appbar, bottomappbar, tab, scaffold, floatingactionbutton, chip, navigationbar, segmentedbutton, badge, slider, divider, bottomsheet, AnimatedBackground |
| 📋 **Pending** | 6 | Foundation Text, NavigationDrawerItem, Overlay, RadioButton, Shadow, navigationrail |
| 🔗 **External widgets** | 12 | button, card, checkbox, dialog, iconbutton, menu, progress, snackbar, switch, text, textfield, tooltip |

### Migration Priority

Based on component importance and usage:

**High Priority (1 component):**
- Foundation Text: Core text component, used throughout all demos and apps

**Medium Priority (5 components):**
- NavigationDrawerItem: Uses hardcoded colors for selected/unselected state
- Overlay: Layout component for modal overlays
- RadioButton: Native component with color options
- Shadow: Modifier for drop shadows
- navigationrail: Expose ContainerColor/ContentColor as ColorDescriptor options

### Component Type Breakdown

**Native Components (19):** Built in go-compose, support or should support ColorDescriptor
- 13 migrated: surface, appbar, bottomappbar, tab, scaffold, floatingactionbutton, chip, navigationbar, segmentedbutton, badge, slider, divider, bottomsheet, AnimatedBackground
- 6 pending: Foundation Text, NavigationDrawerItem, Overlay, RadioButton, Shadow, navigationrail

**External Widget Components (12):** Wrap gio-mw widgets, use gio-mw themes
- No color customization exposed
- Automatically theme-aware via gio-mw

## Notes

- 13 of 19 native components/modifiers now use `theme.ColorDescriptor` (68% complete)
- 6 components pending: Foundation Text, NavigationDrawerItem, Overlay, RadioButton, Shadow, navigationrail
- External widget components (12) use gio-mw themes - no migration needed
- All component statuses verified by code inspection (2025-12-16)
- `SpecificColor()` should only wrap non-theme colors in migrated components
- Theme role selectors ensure proper light/dark theme support
