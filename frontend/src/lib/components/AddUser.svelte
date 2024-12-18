<script lang="ts">
	import { hashWithSHA256 } from '$lib/functions';
	import { userAddedStatus, wantsToAddUser } from '../../stores';
	import AlertMessage from './AlertMessage.svelte';
	import * as AlertMessageConstants from '../../constants/AlertMessageConstants';
	import { onMount } from 'svelte';

	// Lauko kintamieji. Kaip C++ struct
	interface InputData {
		value: string;
		error: boolean;
		errorMessage: string;
	}

	// Nustato default reikšmes laukams
	const getDefaultInputData = (): InputData => ({ value: '', error: false, errorMessage: '' });

	// Laukai
	let name: InputData = getDefaultInputData(),
		lastname: InputData = getDefaultInputData(),
		email: InputData = getDefaultInputData(),
		username: InputData = getDefaultInputData(),
		password: InputData = getDefaultInputData(),
		passwordRepeat: InputData = getDefaultInputData();
	let userType: string;

	function setDefaultInputData() {
		name = getDefaultInputData();
		lastname = getDefaultInputData();
		email = getDefaultInputData();
		username = getDefaultInputData();
		password = getDefaultInputData();
		passwordRepeat = getDefaultInputData();
	}

	// KONSTANTOS

	const CONST_VALID_PASSWORD_LENGTH = 8;
	const CONST_VALID_MIN_USERNAME_LENGTH = 5;
	const CONST_VALID_MAX_USERNAME_LENGTH = 30;
	const CONST_VALID_MAX_NAME_LASTNAME_LENGTH = 50;

	const ERROR_MESSAGE_EMPTY_INPUT = 'Užpildykite lauką.';
	const ERROR_MESSAGE_PASSWORDS_NOT_MATCH = 'Slaptažodžiai nesutampa.';
	const ERROR_MESSAGE_PASSWORD_TOO_SHORT = `Slaptažodis turi būti sudarytas bent iš ${CONST_VALID_PASSWORD_LENGTH} simbolių.`;
	const ERROR_MESSAGE_USERNAME_BAD_LENGTH = `Vartotojo vardas turi būti nuo ${CONST_VALID_MIN_USERNAME_LENGTH} iki ${CONST_VALID_MAX_USERNAME_LENGTH} simbolių.`;
	const ERROR_MESSAGE_EMAIL_INVALID = 'Blogas elektroninio pašto formatas.';
	const ERROR_MESSAGE_NAME_LASTNAME_BAD_LENGTH = `Lauko ilgis turi būti iki ${CONST_VALID_MAX_NAME_LASTNAME_LENGTH} simbolių.`;
	const ERROR_MESSAGE_CONSISTS_NOT_ONLY_LETTERS = 'Laukas turi būti sudarytas tik iš raidžių.';
	const ERROR_MESSAGE_CONSISTS_NOT_ONLY_LETTERS_NUMBERS =
		'Laukas turi būti sudarytas tik iš raidžių ir skaičių.';

	// Funkcija, skirta atnaujinti lauką (kad suveiktų Svelte reactivity (puslapis rerender'intų))
	function refreshInputData(input: InputData) {
		if (input === name) {
			name = { ...input };
		} else if (input === lastname) {
			lastname = { ...input };
		} else if (input === email) {
			email = { ...input };
		} else if (input === username) {
			username = { ...input };
		} else if (input === password) {
			password = { ...input };
		} else if (input === passwordRepeat) {
			passwordRepeat = { ...input };
		}
	}

	// Atnaujina lauko reikšmes
	function updateInputData(input: InputData, changes: Partial<InputData>) {
		Object.assign(input, changes);
	}

	// Prideda error'ą po lauku
	function setInputDataError(input: InputData, message: string) {
		updateInputData(input, { error: true, errorMessage: message });
		refreshInputData(input);
	}

	// Ištrina error'ą iš po lauko
	function removeInputDataError(input: InputData) {
		updateInputData(input, { error: false, errorMessage: '' });
		refreshInputData(input);
	}

	// Gauna registracijos laukų list'ą
	function getRegisterInputs(): InputData[] {
		return [name, lastname, email, username, password, passwordRepeat];
	}

	// Parodo arba panaikina error puslapyje po input laukais
	function updateInputDataErrors() {
		const inputs = getRegisterInputs();

		inputs.forEach((input) => {
			if (!input.value) {
				setInputDataError(input, ERROR_MESSAGE_EMPTY_INPUT);
			} else {
				removeInputDataError(input);
			}
		});

		if (!name.error) {
			if (name.value.length > CONST_VALID_MAX_NAME_LASTNAME_LENGTH) {
				setInputDataError(name, ERROR_MESSAGE_NAME_LASTNAME_BAD_LENGTH);
			} else if (!hasOnlyLetters(name.value)) {
				setInputDataError(name, ERROR_MESSAGE_CONSISTS_NOT_ONLY_LETTERS);
			}
		}

		if (!lastname.error) {
			if (lastname.value.length > CONST_VALID_MAX_NAME_LASTNAME_LENGTH) {
				setInputDataError(lastname, ERROR_MESSAGE_NAME_LASTNAME_BAD_LENGTH);
			} else if (!hasOnlyLetters(lastname.value)) {
				setInputDataError(lastname, ERROR_MESSAGE_CONSISTS_NOT_ONLY_LETTERS);
			}
		}

		if (!email.error) {
			if (!isEmailValid(email.value)) {
				setInputDataError(email, ERROR_MESSAGE_EMAIL_INVALID);
			}
		}

		if (!username.error) {
			if (
				username.value.length < CONST_VALID_MIN_USERNAME_LENGTH ||
				username.value.length > CONST_VALID_MAX_USERNAME_LENGTH
			) {
				setInputDataError(username, ERROR_MESSAGE_USERNAME_BAD_LENGTH);
			} else if (!hasOnlyLettersNumbers(username.value)) {
				setInputDataError(username, ERROR_MESSAGE_CONSISTS_NOT_ONLY_LETTERS_NUMBERS);
			}
		}

		if (!password.error) {
			if (password.value.length < CONST_VALID_PASSWORD_LENGTH) {
				setInputDataError(password, ERROR_MESSAGE_PASSWORD_TOO_SHORT);
			}
		}

		if (!passwordRepeat.error) {
			if (password.value !== passwordRepeat.value) {
				setInputDataError(passwordRepeat, ERROR_MESSAGE_PASSWORDS_NOT_MATCH);
			}
		}
	}

	// Patikrina, ar reikšmė turi tik raides
	function hasOnlyLetters(value: string) {
		const lettersRegex = new RegExp(/^\p{L}+$/u);

		return value.match(lettersRegex);
	}

	// Patikrina, ar reikšmė turi tik raides ir skaičius
	function hasOnlyLettersNumbers(value: string) {
		const lettersNumbersRegex = new RegExp(/^[\p{L}0-9]+$/u);

		return value.match(lettersNumbersRegex);
	}

	// Patikrina, ar elektroninis paštas atitinka formatą
	function isEmailValid(email: string) {
		const emailRegex = new RegExp(
			/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
		);

		return email.match(emailRegex);
	}

	// Ar yra nors viena klaida prisijungime/registracijoje
	function areAnyInputDataErrors(): boolean {
		const inputs = getRegisterInputs();

		return inputs.some((input) => input.error);
	}

	interface AlertMessageHandler {
		show: boolean;
		status: string;
		message: string;
	}

	const alertMessageHandler: AlertMessageHandler = { show: false, status: '', message: '' };
	function showAlertMessage(status: string, message: string) {
		if (alertMessageHandler.show) {
			return;
		}

		alertMessageHandler.show = true;
		alertMessageHandler.status = status;
		alertMessageHandler.message = message;
		setTimeout(() => {
			alertMessageHandler.show = false;
		}, 5000);
	}

	onMount(() => {
		userAddedStatus.set(0);
	});

	async function sendAddNewUserRequest() {
		anyRequestPending = true;

		const response = await fetch('/api/adduser', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({
				username: username.value,
				password: await hashWithSHA256(password.value),
				name: name.value,
				surname: lastname.value,
				email: email.value,
				role: userType
			})
		});
		anyRequestPending = false;

		if (response.status === 200) {
			userAddedStatus.set(200);
			cancelAddProcess();
		} else if (response.status === 409) {
			showAlertMessage(
				AlertMessageConstants.STATUS_FAILURE,
				'Toks vartotojo vardas arba elektroninis paštas jau egzistuoja.'
			);
		} else if (response.status === 500) {
			showAlertMessage(AlertMessageConstants.STATUS_FAILURE, 'Klaida serveryje.');
		} else {
			showAlertMessage(AlertMessageConstants.STATUS_FAILURE, 'Įvyko nežinoma klaida.');
		}
	}

	let anyRequestPending: boolean = false;

	// Pagrindinė request siuntimo funkcija
	function sendRequest() {
		if (anyRequestPending) {
			return;
		}

		updateInputDataErrors();

		if (areAnyInputDataErrors()) {
			return;
		}

		sendAddNewUserRequest();
	}

	function cancelAddProcess() {
		wantsToAddUser.set(false);
	}
</script>

<button
	onclick={cancelAddProcess}
	aria-label="Atšaukti naujo vartotojo pridėjimą"
	id="overlay"
	class="fixed left-0 top-0 z-10 h-full w-full cursor-default bg-black opacity-70"
></button>
<div
	id="modal"
	class="fixed left-1/2 top-1/2 z-10 h-fit w-96 -translate-x-1/2 -translate-y-1/2 transform
	rounded-lg bg-white px-6 pb-3 pt-2"
>
	<p class="mb-5 mt-1 text-2xl font-medium">Naujo vartotojo pridėjimas</p>

	<div class="mb-8 flex flex-col gap-4">
		<div class="flex flex-col">
			<label
				for="name"
				class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
				>Vardas</label
			>
			<input
				id="name"
				bind:value={name.value}
				type="text"
				class="{name.error
					? 'border-red-600'
					: ''} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
			/>
			{#if name.error}
				<p class="ml-1 mt-1 text-xs font-medium text-red-600">
					{name.errorMessage}
				</p>
			{/if}
		</div>
		<div class="flex flex-col">
			<label
				for="lastname"
				class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
				>Pavardė</label
			>
			<input
				id="lastname"
				bind:value={lastname.value}
				type="text"
				class="{lastname.error
					? 'border-red-600'
					: ''} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
			/>
			{#if lastname.error}
				<p class="ml-1 mt-1 text-xs font-medium text-red-600">
					{lastname.errorMessage}
				</p>
			{/if}
		</div>
		<div class="flex flex-col">
			<label
				for="email-register"
				class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
				>Elektroninis paštas</label
			>
			<input
				id="email-register"
				bind:value={email.value}
				type="email"
				class="{email.error
					? 'border-red-600'
					: ''} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
			/>
			{#if email.error}
				<p class="ml-1 mt-1 text-xs font-medium text-red-600">
					{email.errorMessage}
				</p>
			{/if}
		</div>
		<div class="flex flex-col">
			<label
				for="username"
				class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
				>Vartotojo vardas</label
			>
			<input
				id="username"
				bind:value={username.value}
				type="text"
				class="{username.error
					? 'border-red-600'
					: ''} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
			/>
			{#if username.error}
				<p class="ml-1 mt-1 text-xs font-medium text-red-600">
					{username.errorMessage}
				</p>
			{/if}
		</div>
		<div class="flex flex-col">
			<label
				for="password-register"
				class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
				>Slaptažodis</label
			>
			<input
				id="password-register"
				bind:value={password.value}
				type="password"
				class="{password.error
					? 'border-red-600'
					: ''} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
			/>
			{#if password.error}
				<p class="ml-1 mt-1 text-xs font-medium text-red-600">
					{password.errorMessage}
				</p>
			{/if}
		</div>
		<div class="flex flex-col">
			<label
				for="password-repeat"
				class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
				>Pakartokite slaptažodį</label
			>
			<input
				id="password-repeat"
				bind:value={passwordRepeat.value}
				type="password"
				class="{passwordRepeat.error
					? 'border-red-600'
					: ''} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
			/>
			{#if passwordRepeat.error}
				<p class="ml-1 mt-1 text-xs font-medium text-red-600">
					{passwordRepeat.errorMessage}
				</p>
			{/if}
		</div>
		<div class="flex flex-col">
			<label
				for="role"
				class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
				>Vartotojo tipas</label
			>
			<select
				bind:value={userType}
				id="role"
				class="rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
			>
				<option value="customer">Klientas</option>
				<option value="roaster">Darbuotojas</option>
				<option value="admin">Administratorius</option>
			</select>
		</div>
	</div>

	<div class="flex justify-end gap-3">
		<button
			onclick={cancelAddProcess}
			class="10 rounded border border-gray-600 bg-white px-6 py-1 font-medium text-gray-600 hover:bg-gray-200"
			>Atšaukti</button
		>
		<button
			onclick={sendRequest}
			class="rounded bg-green-500 px-6 py-1 font-medium text-white hover:bg-green-600"
			>Sukurti</button
		>
	</div>
</div>

{#if alertMessageHandler.show}
	<AlertMessage status={alertMessageHandler.status} message={alertMessageHandler.message} />
{/if}
