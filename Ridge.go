package main

type Ridge struct {
	W     Matrix
	alpha float64
}

func (model *Ridge) fit(inputs, outputs Matrix) {
	model.W = inputs.transpose().mult(inputs).transpose()
	model.W = eye(model.W.n).scalarMult(model.alpha)
	model.W = model.W.inversed().mult(inputs.transpose()).mult(outputs)
}

func (model *Ridge) predict(input Matrix) Matrix {
	return model.W.mult(input)
}
