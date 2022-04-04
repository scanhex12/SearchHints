package src

type Ridge struct {
	W     Matrix
	alpha float64
}

func (model *Ridge) fit(inputs, outputs Matrix) {
	model.W = inputs.Transpose().Mult(inputs).Transpose()
	model.W = Eye(model.W.N).scalarMult(model.alpha)
	model.W = model.W.Inversed().Mult(inputs.Transpose()).Mult(outputs)
}

func (model *Ridge) predict(input Matrix) Matrix {
	return model.W.Mult(input)
}
