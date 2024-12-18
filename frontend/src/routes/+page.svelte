<script lang="ts">
	import AlertMessage from '../lib/components/AlertMessage.svelte';
	import Footer from '../components/Footer.svelte';
	import * as AlertMessageConstants from '../constants/AlertMessageConstants';
	import { hashWithSHA256 } from '$lib/functions';

	const LOGIN_FORM_TYPE: number = 0,
		REGISTER_FORM_TYPE: number = 1;
	let currentFormType: number = LOGIN_FORM_TYPE;
	const LOGIN_BUTTON_NAME: string = 'Prisijungti',
		REGISTER_BUTTON_NAME: string = 'Registruotis';
	let buttonName: string = LOGIN_BUTTON_NAME;

	// Pakeičia formos tipą (prisijungimas ar registracija)
	function changeFormType(type: number) {
		currentFormType = type;
		buttonName = currentFormType === LOGIN_FORM_TYPE ? LOGIN_BUTTON_NAME : REGISTER_BUTTON_NAME;
	}

	// Lauko kintamieji. Kaip C++ struct
	interface InputData {
		value: string;
		error: boolean;
		errorMessage: string;
	}

	// Nustato default reikšmes laukams
	const getDefaultInputData = (): InputData => ({ value: '', error: false, errorMessage: '' });

	// Login laukai
	let usernameLogin: InputData = getDefaultInputData(),
		passwordLogin: InputData = getDefaultInputData();

	// Register laukai
	let name: InputData = getDefaultInputData(),
		lastname: InputData = getDefaultInputData(),
		email: InputData = getDefaultInputData(),
		usernameRegister: InputData = getDefaultInputData(),
		passwordRegister: InputData = getDefaultInputData(),
		passwordRepeat: InputData = getDefaultInputData();

	function setDefaultInputData(forWhich: number) {
		if (forWhich === LOGIN_FORM_TYPE) {
			usernameLogin = getDefaultInputData();
			passwordLogin = getDefaultInputData();
		} else if (forWhich === REGISTER_FORM_TYPE) {
			name = getDefaultInputData();
			lastname = getDefaultInputData();
			email = getDefaultInputData();
			usernameRegister = getDefaultInputData();
			passwordRegister = getDefaultInputData();
			passwordRepeat = getDefaultInputData();
		}
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
	const ERROR_MESSAGE_USERNAME_EMAIL_EXISTS =
		'Vartotojo vardas arba elektroninis paštas jau egzistuoja.';

	const ERROR_SERVER_USERNAME_EMAIL_EXISTS =
		'This username or email is already taken. Please choose another.';

	// Funkcija, skirta atnaujinti lauką (kad suveiktų Svelte reactivity (puslapis rerender'intų))
	function refreshInputData(input: InputData) {
		if (input === usernameLogin) {
			usernameLogin = { ...input };
		} else if (input === passwordLogin) {
			passwordLogin = { ...input };
		} else if (input === name) {
			name = { ...input };
		} else if (input === lastname) {
			lastname = { ...input };
		} else if (input === email) {
			email = { ...input };
		} else if (input === usernameRegister) {
			usernameRegister = { ...input };
		} else if (input === passwordRegister) {
			passwordRegister = { ...input };
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

	// Gauna prisijungimo laukų list'ą
	function getLoginInputs(): InputData[] {
		return [usernameLogin, passwordLogin];
	}

	// Gauna registracijos laukų list'ą
	function getRegisterInputs(): InputData[] {
		return [name, lastname, email, usernameRegister, passwordRegister, passwordRepeat];
	}

	// Parodo arba panaikina error puslapyje po input laukais
	function updateInputDataErrors() {
		const inputs = currentFormType === LOGIN_FORM_TYPE ? getLoginInputs() : getRegisterInputs();

		inputs.forEach((input) => {
			if (!input.value) {
				setInputDataError(input, ERROR_MESSAGE_EMPTY_INPUT);
			} else {
				removeInputDataError(input);
			}
		});

		if (currentFormType === REGISTER_FORM_TYPE) {
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

			if (!usernameRegister.error) {
				if (
					usernameRegister.value.length < CONST_VALID_MIN_USERNAME_LENGTH ||
					usernameRegister.value.length > CONST_VALID_MAX_USERNAME_LENGTH
				) {
					setInputDataError(usernameRegister, ERROR_MESSAGE_USERNAME_BAD_LENGTH);
				} else if (!hasOnlyLettersNumbers(usernameRegister.value)) {
					setInputDataError(usernameRegister, ERROR_MESSAGE_CONSISTS_NOT_ONLY_LETTERS_NUMBERS);
				}
			}

			if (!passwordRegister.error) {
				if (passwordRegister.value.length < CONST_VALID_PASSWORD_LENGTH) {
					setInputDataError(passwordRegister, ERROR_MESSAGE_PASSWORD_TOO_SHORT);
				}
			}

			if (!passwordRepeat.error) {
				if (passwordRegister.value !== passwordRepeat.value) {
					setInputDataError(passwordRepeat, ERROR_MESSAGE_PASSWORDS_NOT_MATCH);
				}
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
		const inputs = currentFormType === LOGIN_FORM_TYPE ? getLoginInputs() : getRegisterInputs();

		return inputs.some((input) => input.error);
	}

	let anyRequestPending: boolean = false;

	// Login request
	async function sendLoginRequest() {
		anyRequestPending = true;
		const response = await fetch('/api/login', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({
				username: usernameLogin.value,
				password: passwordLogin.value
			})
		});
		anyRequestPending = false;

		if (response.status === 200) {
			showAlertMessage(AlertMessageConstants.STATUS_SUCCESS, 'Prisijungimas sėkmingas!');

			const params = new URLSearchParams(window.location.search);
			if (params.has('redirectTo')) {
				window.location.assign(params.get('redirectTo'));
			} else {
				window.location.assign('/main');
			}
		} else if (response.status === 401) {
			showAlertMessage(AlertMessageConstants.STATUS_FAILURE, 'Blogi prisijungimo duomenys.');
		} else if (response.status === 500) {
			showAlertMessage(AlertMessageConstants.STATUS_FAILURE, 'Klaida serveryje.');
		} else {
			showAlertMessage(AlertMessageConstants.STATUS_FAILURE, 'Įvyko nežinoma klaida.');
		}
	}

	// Register request
	async function sendRegisterRequest() {
		anyRequestPending = true;

		const response = await fetch('/api/register', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({
				username: usernameRegister.value,
				password: await hashWithSHA256(passwordRegister.value),
				name: name.value,
				surname: lastname.value,
				email: email.value,
				role: 'customer'
			})
		});
		anyRequestPending = false;

		if (response.status === 200) {
			showAlertMessage(AlertMessageConstants.STATUS_SUCCESS, 'Registracija sėkminga!');
			currentFormType = LOGIN_FORM_TYPE;
			buttonName = LOGIN_BUTTON_NAME;
			setDefaultInputData(REGISTER_FORM_TYPE);
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

	// Pagrindinė request siuntimo funkcija
	function sendRequest() {
		if (anyRequestPending) {
			return;
		}

		updateInputDataErrors();

		if (areAnyInputDataErrors()) {
			return;
		}

		if (currentFormType === LOGIN_FORM_TYPE) {
			sendLoginRequest();
		} else {
			sendRegisterRequest();
		}
	}
</script>

<!-- https://svelte.dev/docs/svelte/snippet -->

<!-- Magija -->

<svelte:head>
	<title>Kavos Skrudykla</title>
</svelte:head>

<div class="flex min-h-screen flex-col justify-center">
	<article class="grid grow grid-cols-1 lg:grid-cols-2">
		<div id="right-side" class="flex flex-col pl-5 pt-5 lg:pl-12 lg:pt-32 xl:pl-24 2xl:pl-52">
			<div class="self-center text-6xl font-bold sm:self-auto sm:text-8xl">
				<p>Kavos</p>
				<div class="relative w-fit">
					<p>Skrudykla</p>
					<div class="group absolute right-0 mt-6 h-2 w-1/2 bg-black">
						<span aria-hidden="true" class="bg-page animate-slide absolute inset-0"></span>
					</div>
				</div>
			</div>
			<div class="mt-16 pr-5 text-lg font-medium lg:mt-24 lg:pl-0 lg:pr-12">
				<p>Sveiki atvykę į Kavos Skrudyklą!</p>
				<p>
					Sistema pasiekiama prisijungus su vartotojo paskyra. Jeigu neturite vartotojo paskyros,
					tiesiog užsiregistruokite.
				</p>
				<p>Prisijunkite ir nusipirkite skaniausią skrudintą kavą Lietuvoje !</p>
			</div>
		</div>

		<div id="left-side" class="flex justify-center p-5 lg:justify-normal">
			<div id="form" class="z-10 grow content-center lg:max-w-xl">
				<div id="form-upper" class="grid h-8 grid-cols-2 text-center text-sm">
					<div
						class="bg-main flex cursor-pointer rounded-tl-full rounded-tr-full {currentFormType ===
						LOGIN_FORM_TYPE
							? 'border-l border-r border-t font-medium'
							: 'hover:bg-hover-button border-b'} border-black"
					>
						<button
							on:click={() => changeFormType(LOGIN_FORM_TYPE)}
							class="grow rounded-tl-full rounded-tr-full">Prisijungimas</button
						>
					</div>
					<div
						class="bg-main flex cursor-pointer rounded-tl-full rounded-tr-full {currentFormType ===
						REGISTER_FORM_TYPE
							? 'border-l border-r border-t font-medium'
							: 'hover:bg-hover-button border-b'} border-black"
					>
						<button
							on:click={() => changeFormType(REGISTER_FORM_TYPE)}
							class="grow rounded-tl-full rounded-tr-full">Registracija</button
						>
					</div>
				</div>

				<div
					id="form-main"
					class="bg-main h-fit border-b border-l border-r border-black p-7 pt-5 shadow-sm shadow-black"
				>
					{#if currentFormType === LOGIN_FORM_TYPE}
						<p class="mt-1 px-2 text-xs text-gray-500">
							- Įveskite elektroninį paštą ir slaptažodį norint prisijungti.
						</p>
					{:else}
						<p class="mt-1 px-2 text-xs text-gray-500">
							- Įveskite informaciją norint užsiregistruoti.
						</p>
					{/if}
					{#if currentFormType === LOGIN_FORM_TYPE}
						<div id="form-login" class="my-5 flex flex-col gap-4">
							<div class="flex flex-col">
								<label
									for="email-login"
									class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
									>Vartotojo vardas</label
								>
								<input
									id="email-login"
									bind:value={usernameLogin.value}
									type="email"
									class="{usernameLogin.error
										? 'border-red-600'
										: 'border-button'} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
								/>
								{#if usernameLogin.error}
									<p class="ml-1 mt-1 text-xs font-medium text-red-600">
										{usernameLogin.errorMessage}
									</p>
								{/if}
							</div>
							<div class="flex flex-col">
								<label
									for="password-login"
									class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
									>Slaptažodis</label
								>
								<input
									id="password-login"
									bind:value={passwordLogin.value}
									type="password"
									class="{passwordLogin.error
										? 'border-red-600'
										: 'border-button'} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
								/>
								{#if passwordLogin.error}
									<p class="ml-1 mt-1 text-xs font-medium text-red-600">
										{passwordLogin.errorMessage}
									</p>
								{/if}
							</div>
						</div>
					{/if}
					{#if currentFormType === REGISTER_FORM_TYPE}
						<div id="form-register" class="my-5 flex flex-col gap-4">
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
										: 'border-button'} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
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
										: 'border-button'} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
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
										: 'border-button'} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
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
									bind:value={usernameRegister.value}
									type="text"
									class="{usernameRegister.error
										? 'border-red-600'
										: 'border-button'} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
								/>
								{#if usernameRegister.error}
									<p class="ml-1 mt-1 text-xs font-medium text-red-600">
										{usernameRegister.errorMessage}
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
									bind:value={passwordRegister.value}
									type="password"
									class="{passwordRegister.error
										? 'border-red-600'
										: 'border-button'} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
								/>
								{#if passwordRegister.error}
									<p class="ml-1 mt-1 text-xs font-medium text-red-600">
										{passwordRegister.errorMessage}
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
										: 'border-button'} rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
								/>
								{#if passwordRepeat.error}
									<p class="ml-1 mt-1 text-xs font-medium text-red-600">
										{passwordRepeat.errorMessage}
									</p>
								{/if}
							</div>
						</div>
					{/if}
				</div>

				<div
					id="form-lower"
					class="bg-button group relative flex h-10 cursor-pointer gap-2 overflow-hidden rounded-b-lg border-b border-l border-r border-black shadow-sm shadow-black"
				>
					<button
						on:click={() => sendRequest()}
						class="z-10 grow text-end font-semibold tracking-widest"
						>{buttonName.toUpperCase()}
						<span class="mr-2 text-lg">&rarr;</span></button
					>
					<span
						aria-hidden="true"
						class="bg-hover-button -500 absolute inset-0 -translate-x-full transition group-hover:translate-x-0"
					></span>
				</div>
			</div>
		</div>
	</article>

	<Footer />
</div>

{#if alertMessageHandler.show}
	<AlertMessage status={alertMessageHandler.status} message={alertMessageHandler.message} />
{/if}
