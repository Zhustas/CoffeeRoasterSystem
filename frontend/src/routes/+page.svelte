<script lang="ts">
	import AlertMessage from '../lib/components/AlertMessage.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import * as AlertMessageConstants from '$lib/constants/AlertMessageConstants';
	import { alertStore, type AlertMessageHandler, showAlertMessage } from '$lib/alertmessagehandler';
	import * as FormTypeConstants from '$lib/constants/FormTypeConstants';
	import * as InputConstants from '$lib/constants/InputConstants';
	import type { InputData } from '$lib/inputhandler';
	import * as InputHandler from '$lib/inputhandler';
	import { hashWithSHA256 } from '$lib/functions';
	import { onMount } from 'svelte';

	onMount(() => {
		// const search = new URLSearchParams(window.location.search);
		// if (search.has('type') && search.get('type') === 'register') {
		// 	currentFormType = FormTypeConstants.FORM_TYPE_REGISTER;
		// }
	});

	let alertMessageHandler: AlertMessageHandler = $state();
	alertStore.subscribe((state) => {
		alertMessageHandler = state;
	});

	let currentFormType = $state(FormTypeConstants.FORM_TYPE_LOGIN);

	const LOGIN_BUTTON_NAME: string = 'Prisijungti';
	const REGISTER_BUTTON_NAME: string = 'Registruotis';
	let buttonName: string = $derived(currentFormType === FormTypeConstants.FORM_TYPE_LOGIN ? LOGIN_BUTTON_NAME : REGISTER_BUTTON_NAME);

	let usernameLogin: InputData = $state(InputHandler.getDefaultInputData()),
		passwordLogin: InputData = $state(InputHandler.getDefaultInputData());

	let name: InputData = $state(InputHandler.getDefaultInputData()),
		lastname: InputData = $state(InputHandler.getDefaultInputData()),
		email: InputData = $state(InputHandler.getDefaultInputData()),
		usernameRegister: InputData = $state(InputHandler.getDefaultInputData()),
		passwordRegister: InputData = $state(InputHandler.getDefaultInputData()),
		passwordRepeat: InputData = $state(InputHandler.getDefaultInputData());

	function changeFormType(type: any) {
		currentFormType = type;
	}

	function getLoginInputs(): InputData[] {
		return [usernameLogin, passwordLogin];
	}

	function getRegisterInputs(): InputData[] {
		return [name, lastname, email, usernameRegister, passwordRegister, passwordRepeat];
	}

	function updateInputDataErrors() {
		const inputs = currentFormType === FormTypeConstants.FORM_TYPE_LOGIN ? getLoginInputs() : getRegisterInputs();
		inputs.forEach((input) => {
			if (!input.value) {
				InputHandler.setInputDataError(input, InputConstants.ERROR_EMPTY_INPUT);
			} else {
				InputHandler.removeInputDataError(input);
			}
		});

		if (currentFormType === FormTypeConstants.FORM_TYPE_REGISTER) {
			if (!name.error) {
				if (name.value.length > InputConstants.VALID_MAX_NAME_LENGTH) {
					InputHandler.setInputDataError(name, InputConstants.ERROR_NAME_BAD_LENGTH);
				} else if (!hasOnlyLetters(name.value)) {
					InputHandler.setInputDataError(name, InputConstants.ERROR_CONSISTS_NOT_ONLY_LETTERS);
				}
			}

			if (!lastname.error) {
				if (lastname.value.length > InputConstants.VALID_MAX_LASTNAME_LENGTH) {
					InputHandler.setInputDataError(lastname, InputConstants.ERROR_LASTNAME_BAD_LENGTH);
				} else if (!hasOnlyLetters(lastname.value)) {
					InputHandler.setInputDataError(lastname, InputConstants.ERROR_CONSISTS_NOT_ONLY_LETTERS);
				}
			}

			if (!email.error) {
				if (!isEmailValid(email.value)) {
					InputHandler.setInputDataError(email, InputConstants.ERROR_EMAIL_INVALID);
				}
			}

			if (!usernameRegister.error) {
				if (
					usernameRegister.value.length < InputConstants.VALID_MIN_USERNAME_LENGTH ||
					usernameRegister.value.length > InputConstants.VALID_MAX_USERNAME_LENGTH
				) {
					InputHandler.setInputDataError(usernameRegister, InputConstants.ERROR_USERNAME_BAD_LENGTH);
				} else if (!hasOnlyLettersNumbers(usernameRegister.value)) {
					InputHandler.setInputDataError(usernameRegister, InputConstants.ERROR_CONSISTS_NOT_ONLY_LETTERS_NUMBERS);
				}
			}

			if (!passwordRegister.error) {
				if (passwordRegister.value.length < InputConstants.VALID_PASSWORD_LENGTH) {
					InputHandler.setInputDataError(passwordRegister, InputConstants.ERROR_PASSWORD_TOO_SHORT);
				}
			}

			if (!passwordRepeat.error) {
				if (passwordRegister.value !== passwordRepeat.value) {
					InputHandler.setInputDataError(passwordRepeat, InputConstants.ERROR_PASSWORDS_NOT_MATCH);
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
		const inputs = currentFormType === FormTypeConstants.FORM_TYPE_LOGIN ? getLoginInputs() : getRegisterInputs();

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
			currentFormType = FormTypeConstants.FORM_TYPE_LOGIN;
			InputHandler.setDefaultInputData(getRegisterInputs());
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

	// Pagrindinė request siuntimo funkcija
	function sendRequest() {
		if (anyRequestPending) {
			return;
		}

		updateInputDataErrors();

		if (areAnyInputDataErrors()) {
			return;
		}

		if (currentFormType === FormTypeConstants.FORM_TYPE_LOGIN) {
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
						<span aria-hidden="true" class="absolute inset-0 animate-slide bg-page"></span>
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
						class="flex cursor-pointer rounded-tl-full rounded-tr-full bg-main {currentFormType ===
						FormTypeConstants.FORM_TYPE_LOGIN
							? 'border-l border-r border-t font-medium'
							: 'border-b hover:bg-hover-button'} border-black"
					>
						<button
							onclick={() => changeFormType(FormTypeConstants.FORM_TYPE_LOGIN)}
							class="grow rounded-tl-full rounded-tr-full">Prisijungimas
						</button
						>
					</div>
					<div
						class="flex cursor-pointer rounded-tl-full rounded-tr-full bg-main {currentFormType ===
						FormTypeConstants.FORM_TYPE_REGISTER
							? 'border-l border-r border-t font-medium'
							: 'border-b hover:bg-hover-button'} border-black"
					>
						<button
							onclick={() => changeFormType(FormTypeConstants.FORM_TYPE_REGISTER)}
							class="grow rounded-tl-full rounded-tr-full">Registracija
						</button
						>
					</div>
				</div>

				<div
					id="form-main"
					class="h-fit border-b border-l border-r border-black bg-main p-7 pt-5 shadow-sm shadow-black"
				>
					{#if currentFormType === FormTypeConstants.FORM_TYPE_LOGIN}
						<p class="mt-1 px-2 text-xs text-gray-500">
							- Įveskite vartotojo vardą ir slaptažodį norint prisijungti.
						</p>
					{:else}
						<p class="mt-1 px-2 text-xs text-gray-500">
							- Įveskite informaciją norint užsiregistruoti.
						</p>
					{/if}
					{#if currentFormType === FormTypeConstants.FORM_TYPE_LOGIN}
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
									type="text"
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
					{#if currentFormType === FormTypeConstants.FORM_TYPE_REGISTER}
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
					class="group relative flex h-10 cursor-pointer gap-2 overflow-hidden rounded-b-lg border-b border-l border-r border-black bg-button shadow-sm shadow-black"
				>
					<button
						onclick={() => sendRequest()}
						class="z-10 grow text-end font-semibold tracking-widest"
					>{buttonName.toUpperCase()}
						<span class="mr-2 text-lg">&rarr;</span></button
					>
					<span
						aria-hidden="true"
						class="-500 absolute inset-0 -translate-x-full bg-hover-button transition group-hover:translate-x-0"
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
