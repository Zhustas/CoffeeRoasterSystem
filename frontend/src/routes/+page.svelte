<script lang="ts">
	import AlertMessage from '../components/AlertMessage.svelte';
	import * as AlertMessageConstants from '../constants/AlertMessageConstants';

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

	// Error'ų tipai
	const ERROR_MESSAGE_EMPTY_INPUT = 'Užpildykite lauką.';
	const ERROR_MESSAGE_PASSWORDS_NOT_MATCH = 'Slaptažodžiai nesutampa.';

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

		if (currentFormType === REGISTER_FORM_TYPE) {
			if (passwordRegister.value !== passwordRepeat.value) {
				setInputDataError(passwordRepeat, ERROR_MESSAGE_PASSWORDS_NOT_MATCH);
			}
		}

		inputs.forEach((input) => {
			if (!input.value) {
				setInputDataError(input, ERROR_MESSAGE_EMPTY_INPUT);
			} else {
				removeInputDataError(input);
			}
		});
	}

	// Ar yra nors viena klaida prisijungime/registracijoje
	function areAnyInputDataErrors(): boolean {
		const inputs = currentFormType === LOGIN_FORM_TYPE ? getLoginInputs() : getRegisterInputs();

		return inputs.some((input) => input.error);
	}

	let anyRequestPending: boolean = false;
	let showMessage: boolean = false;

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
			window.location.assign('/main');
		}
		console.log(response.status);
	}

	// Register request
	async function sendRegisterRequest() {
		anyRequestPending = true;
		const response = await fetch('/api/register', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({
				username: usernameRegister.value,
				password: passwordRegister.value,
				name: name.value,
				surname: lastname.value,
				email: email.value,
				role: 'customer'
			})
		});
		anyRequestPending = false;

		if (response.status === 200) {
			showMessage = true;
			setTimeout(() => {
				showMessage = false;
			}, 5000);
			currentFormType = LOGIN_FORM_TYPE;
			buttonName = LOGIN_BUTTON_NAME;
			setDefaultInputData(REGISTER_FORM_TYPE);
		} else {
			console.log(await response.json());
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

		if (currentFormType === LOGIN_FORM_TYPE) {
			sendLoginRequest();
		} else {
			sendRegisterRequest();
		}
	}
</script>

<!-- https://svelte.dev/docs/svelte/snippet -->

<!-- Magija -->
<div class="flex h-screen flex-col">
	<article class="grid grow grid-cols-2">
		<div id="right-side" class="flex flex-col pl-52 pt-32">
			<div class="text-8xl font-bold">
				<p>Kavos</p>
				<div class="relative w-fit">
					<p>Skrudykla</p>
					<div class="group absolute right-0 mt-6 h-2 w-1/2 bg-black">
						<span aria-hidden="true" class="bg-page animate-slide absolute inset-0"></span>
					</div>
				</div>
			</div>
			<div class="mt-24 pr-12 text-lg font-medium">
				<p>Sveiki atvykę į Kavos Skrudyklą!</p>
				<p>
					Sistema pasiekiama prisijungus su vartotojo paskyra. Jeigu neturite vartotojo paskyros,
					tiesiog užsiregistruokite.
				</p>
				<p>Prisijunkite ir nusipirkite skaniausią skrudintą kavą Lietuvoje !</p>
			</div>
		</div>

		<div id="left-side" class="flex p-5">
			<div id="form" class="max-w-xl grow content-center">
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

	<footer class="flex h-12 border-t border-gray-400 bg-white">
		<img alt="Kavos Skrudykla logo" src="logo.png" class="ml-5 h-7 w-7 self-center rounded-full" />
		<p class="ml-5 self-center text-xs font-medium text-gray-400">
			&copy; 2024 UAB "Kavos Skrudykla"
		</p>
	</footer>
</div>

{#if showMessage}
	<AlertMessage status={AlertMessageConstants.STATUS_SUCCESS} message="Registracija sėkminga!" />
{/if}
