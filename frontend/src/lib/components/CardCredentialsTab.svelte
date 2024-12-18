<script lang="ts">
	import { wantsToMakeOrder } from '../../stores';

	const { makeOrder } = $props();

	// Lauko kintamieji. Kaip C++ struct
	interface InputData {
		value: string;
		error: boolean;
		errorMessage: string;
	}

	// Nustato default reikšmes laukams
	const getDefaultInputData = (): InputData => ({ value: '', error: false, errorMessage: '' });

	// Funkcija, skirta atnaujinti lauką (kad suveiktų Svelte reactivity (puslapis rerender'intų))
	function refreshInputData(input: InputData) {
		if (input === cardNumber) {
			cardNumber = { ...input };
		} else if (input === cardName) {
			cardName = { ...input };
		} else if (input === cardExpiration) {
			cardExpiration = { ...input };
		} else if (input === cardCVC) {
			cardCVC = { ...input };
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

	const VALID_CARD_NUMBER_LENGTH = 16;

	const ERROR_MESSAGE_EMPTY_INPUT = 'Užpildykite lauką.';
	const ERROR_MESSAGE_CARD_NUMBER_BAD_LENGTH = `Lauko ilgis turi būti ${VALID_CARD_NUMBER_LENGTH} simbolių.`;
	const ERROR_MESSAGE_NOT_ONLY_NUMBERS = `Laukas turi būti sudarytas iš skaičių.`;
	const ERROR_MESSAGE_NOT_ONLY_LETTERS = `Laukas turi būti sudarytas iš raidžių.`;

	// Patikrina, ar reikšmė turi tik raides
	function hasOnlyNumbers(value: string) {
		const lettersRegex = new RegExp(/^[0-9]*$/u);

		return value.match(lettersRegex);
	}

	// Patikrina, ar reikšmė turi tik raides
	function hasOnlyLetters(value: string) {
		const lettersRegex = new RegExp(/^\p{L}+$/u);

		return value.match(lettersRegex);
	}

	// Parodo arba panaikina error puslapyje po input laukais
	function updateInputDataErrors() {
		const inputs = getInputs();

		inputs.forEach((input) => {
			if (!input.value) {
				setInputDataError(input, ERROR_MESSAGE_EMPTY_INPUT);
			} else {
				removeInputDataError(input);
			}
		});

		if (!cardNumber.error) {
			if (cardNumber.value.length !== VALID_CARD_NUMBER_LENGTH) {
				setInputDataError(cardNumber, ERROR_MESSAGE_CARD_NUMBER_BAD_LENGTH);
			} else if (!hasOnlyNumbers(cardNumber.value)) {
				setInputDataError(cardNumber, ERROR_MESSAGE_NOT_ONLY_NUMBERS);
			}
		}

		// if (!cardName.error) {
		// 	if (!hasOnlyLetters(cardName.value)) {
		// 		setInputDataError(cardName, ERROR_MESSAGE_NOT_ONLY_LETTERS);
		// 	}
		// }

		if (!cardCVC.error) {
			if (!hasOnlyNumbers(cardCVC.value)) {
				setInputDataError(cardCVC, ERROR_MESSAGE_NOT_ONLY_NUMBERS);
			}
		}
	}

	// Gauna prisijungimo laukų list'ą
	function getInputs(): InputData[] {
		return [cardNumber, cardName, cardExpiration, cardCVC];
	}

	// Ar yra nors viena klaida prisijungime/registracijoje
	function areAnyInputDataErrors(): boolean {
		const inputs = getInputs();

		return inputs.some((input) => input.error);
	}

	// Pagrindinė request siuntimo funkcija
	function sendRequest() {
		updateInputDataErrors();

		if (areAnyInputDataErrors()) {
			return;
		}

		makeOrder();
		cancelOrdering();
	}

	function cancelOrdering() {
		wantsToMakeOrder.set(false);
	}

	let cardNumber: InputData = $state(getDefaultInputData()),
		cardName: InputData = $state(getDefaultInputData()),
		cardExpiration: InputData = $state(getDefaultInputData()),
		cardCVC: InputData = $state(getDefaultInputData());
</script>

<button
	onclick={cancelOrdering}
	aria-label="Paspauskite, kad baigtumėte naujos kavos kūrimo procesą."
	id="overlay"
	class="fixed left-0 top-0 h-full w-full cursor-default bg-black opacity-70"
></button>
<div
	id="modal"
	class="fixed left-1/2 top-1/2 h-fit w-[28rem] -translate-x-1/2 -translate-y-1/2 transform
	rounded-lg bg-white px-6 pb-3 pt-2"
>
	<p class="mb-4 mt-1 text-2xl font-medium">Apmokėjimas</p>

	<div class="mb-8 flex flex-col">
		<div class="mb-4 flex flex-col">
			<label
				for="cardnumber"
				class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
				>Kortelės numeris</label
			>
			<input
				id="cardnumber"
				bind:value={cardNumber.value}
				type="text"
				class="{cardNumber.error
					? 'border-red-600'
					: 'border-gray-400'} rounded-md hover:border-gray-500 focus:border-black focus:outline-none focus:ring-0"
			/>
			{#if cardNumber.error}
				<p class="ml-1 mt-1 text-xs font-medium text-red-600">
					{cardNumber.errorMessage}
				</p>
			{/if}
		</div>

		<div class="mb-4 flex flex-col">
			<label
				for="cardname"
				class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
				>Vardas ant kortelės</label
			>
			<input
				id="cardname"
				bind:value={cardName.value}
				type="text"
				class="{cardName.error
					? 'border-red-600'
					: 'border-gray-400'} rounded-md hover:border-gray-500 focus:border-black focus:outline-none focus:ring-0"
			/>
			{#if cardName.error}
				<p class="ml-1 mt-1 text-xs font-medium text-red-600">
					{cardName.errorMessage}
				</p>
			{/if}
		</div>

		<div class="mb-4 flex gap-10">
			<div class="flex flex-col">
				<label
					for="cardexpiration"
					class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
					>Galioja iki</label
				>
				<input
					id="cardexpiration"
					bind:value={cardExpiration.value}
					class="{cardExpiration.error
						? 'border-red-600'
						: 'border-gray-400'} w-24 rounded-md hover:border-gray-500 focus:border-black focus:outline-none focus:ring-0"
				/>
				{#if cardExpiration.error}
					<p class="ml-1 mt-1 text-xs font-medium text-red-600">
						{cardExpiration.errorMessage}
					</p>
				{/if}
			</div>

			<div class="flex flex-col">
				<label
					for="cardcvc"
					class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
					>CVC</label
				>
				<input
					id="cardcvc"
					maxlength="3"
					bind:value={cardCVC.value}
					class="{cardCVC.error
						? 'border-red-600'
						: 'border-gray-400'} w-24 rounded-md hover:border-gray-500 focus:border-black focus:outline-none focus:ring-0"
				/>
				{#if cardCVC.error}
					<p class="ml-1 mt-1 text-xs font-medium text-red-600">
						{cardCVC.errorMessage}
					</p>
				{/if}
			</div>
		</div>
	</div>

	<div class="flex justify-end gap-3">
		<button
			onclick={cancelOrdering}
			class="10 rounded border border-gray-600 bg-white px-6 py-1 font-medium text-gray-600 hover:bg-gray-200"
			>Atšaukti</button
		>
		<button
			onclick={sendRequest}
			class="rounded bg-green-500 px-6 py-1 font-medium text-white hover:bg-green-600"
			>Užsakyti</button
		>
	</div>
</div>
