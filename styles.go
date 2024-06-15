package main

import (
	"github.com/chasefleming/elem-go/styles"
)

var (
	primaryWidth     = styles.Pixels(900)
	navWidth         = styles.Pixels(150)
	borderColor      = "White"
	buttonBackground = "#010"
)

var StyleMgr = styles.NewStyleManager()

// TODO: combine media query classes now that I'm getting the hang of it
var NavMediaQuery = styles.CompositeStyle{
	MediaQueries: map[string]styles.Props{
		"@media (min-width: 1275px) ": {
			"float":      "left",
			"width":      navWidth,
			"position":   "fixed",
			"text-align": "left",
			"font-size":  "large",
		},
	},
}

// TODO: the mediaquery part of this doesn't work, however the ::before does
var NavBeforeMediaQuery = styles.CompositeStyle{
	PseudoClasses: map[string]styles.Props{
		"::before": {
			styles.Content:        "\"Navigation\"",
			styles.TextAlign:      "center",
			styles.Display:        "block",
			styles.FontSize:       "large",
			styles.Color:          "white",
			styles.TextDecoration: "underline",
			styles.MarginTop:      styles.Em(0.5),
		},
	}, // missing ::before as constant
	MediaQueries: map[string]styles.Props{
		"@media (min-width: 1275px) ": {
			styles.Content:        "\"Navigation\"",
			styles.TextAlign:      "center",
			styles.Display:        "block",
			styles.FontSize:       "large",
			styles.Color:          "white",
			styles.TextDecoration: "underline",
			styles.MarginTop:      styles.Em(0.5),
		},
	},
}

var NavLiMediaQuery = styles.CompositeStyle{
	MediaQueries: map[string]styles.Props{
		"@media (min-width: 1275px) ": {
			"display":    "block",
			"text-align": "center",
			"margin":     "0.5em auto",
		},
	},
}

var BodyStyle = styles.Props{
	styles.Color:      "White",
	styles.Background: "#011",
}

var HeaderH1Style = styles.Merge(BaseH1Style, styles.Props{
	styles.MaxWidth:      primaryWidth,
	styles.FontSize:      "xxx-large",
	styles.MarginLeft:    "auto",
	styles.MarginRight:   "auto",
	styles.FontVariant:   "small-caps",
	styles.LetterSpacing: "0.5rem",
})

var HeaderH3Style = styles.Props{
	styles.MaxWidth:    primaryWidth,
	styles.FontSize:    "x-large",
	styles.MarginLeft:  "auto",
	styles.MarginRight: "auto",
	styles.FontVariant: "small-caps",
	styles.TextAlign:   "center",
}

var NavStyle = styles.Props{
	styles.MaxWidth:     primaryWidth,
	styles.TextAlign:    "center",
	"clear":             "both", // add as official constant
	styles.Margin:       "auto",
	styles.Border:       "solid 1px " + borderColor,
	styles.BorderRadius: "10px",
	styles.FontVariant:  "small-caps",
	styles.MarginBottom: "1em",
}

var NavAStyle = styles.Props{
	styles.TextDecoration: "none",
	styles.Color:          "white",
}

var NavLiStyle = styles.Props{
	styles.Display:      "inline-block",
	styles.ListStyle:    "none",
	styles.Border:       "dashed 1px " + borderColor,
	styles.BorderRadius: "10px",
	styles.MaxWidth:     "7em",
	styles.Background:   buttonBackground,
}

var NavUlStyle = styles.Props{
	styles.Padding: "0",
	styles.Margin:  "0.5em",
}

var NavAHoverLiClass = styles.CompositeStyle{
	Default:       NavAStyle,
	PseudoClasses: map[string]styles.Props{styles.PseudoHover: NavAHoverLiStyle},
}

var NavAHoverLiStyle = styles.Props{
	styles.Background: "white",
	styles.Color:      "#010",
	styles.BoxShadow:  "2px 2px orange",
}

var FullNavAStyle = styles.Merge(NavAStyle, NavAHoverLiStyle)

var BaseH1Style = styles.Props{
	styles.TextAlign: "center",
}

var BaseH2Style = styles.Props{
	styles.Color: "orange",
}

var BaseH3Style = styles.Props{
	styles.Color: "#ccc",
}

var MainStyle = styles.Props{
	styles.MaxWidth:     primaryWidth,
	styles.Margin:       "auto",
	styles.Border:       "solid 2px " + borderColor,
	styles.Padding:      "1em",
	styles.BorderRadius: "10px",
}
