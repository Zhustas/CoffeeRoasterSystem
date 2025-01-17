interface InputData {
	value: string;
	error: boolean;
	errorMessage: string;
}

const getDefaultInputData = (): InputData => ({ value: '', error: false, errorMessage: '' });

function setDefaultInputData(inputs: InputData[]) {
	inputs.forEach((input) => {
		Object.assign<InputData, Required<InputData>>(input, {
			value: '',
			error: false,
			errorMessage: ''
		});
	});
}

function setInputDataError(input: InputData, message: string) {
	Object.assign<InputData, Partial<InputData>>(input, { error: true, errorMessage: message });
}

function removeInputDataError(input: InputData) {
	Object.assign<InputData, Partial<InputData>>(input, { error: false, errorMessage: '' });
}

export {
	type InputData,
	getDefaultInputData,
	setDefaultInputData,
	setInputDataError,
	removeInputDataError
};
