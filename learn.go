package main

import "neural_network/form"

type Learn struct {
    entity         form.Form
    hideCollection map[string]form.HideNeuron
    state          map[string]float64
    resultWeight   float64
}


