package intl

// https://cs.android.com/androidx/platform/frameworks/support/+/androidx-main:compose/ui/ui-text/src/commonMain/kotlin/androidx/compose/ui/text/intl/LocaleList.kt;drc=4970f6e96cdb06089723da0ab8ec93ae3f067c7a;l=31

type LocaleList struct {
	List []Locale
}

func EmptyLocaleList() LocaleList {
	return LocaleList{
		List: []Locale{},
	}
}
