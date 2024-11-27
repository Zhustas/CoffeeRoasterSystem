<script lang="ts">
	import { wantsToAddCoffee } from '../../stores';

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
		description: InputData = getDefaultInputData(),
		stock: InputData = getDefaultInputData(),
		price: InputData = getDefaultInputData();
	let typeValue;

	const ERROR_MESSAGE_EMPTY_INPUT = 'Užpildykite lauką.';

	// Funkcija, skirta atnaujinti lauką (kad suveiktų Svelte reactivity (puslapis rerender'intų))
	function refreshInputData(input: InputData) {
		if (input === name) {
			name = { ...input };
		} else if (input === description) {
			description = { ...input };
		} else if (input === stock) {
			stock = { ...input };
		} else if (input === price) {
			price = { ...input };
		}
	}

	// Atnaujina lauko reikšmes
	function updateInputData(input: InputData, changes: Partial<InputData>) {
		Object.assign(input, changes);
	}

	// Gauna prisijungimo laukų list'ą
	function getInputs(): InputData[] {
		return [name, description, stock, price];
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
	}

	// Ar yra nors viena klaida prisijungime/registracijoje
	function areAnyInputDataErrors(): boolean {
		const inputs = getInputs();

		return inputs.some((input) => input.error);
	}

	async function sendCreateRequest() {
		const response = await fetch(`/api/addcoffee`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({
				name: name.value,
				roast_type: typeValue,
				description: description.value,
				stock: stock.value,
				price: price.value
			})
		});

		console.log(await response.json());
		cancelCreate();
	}

	// Pagrindinė request siuntimo funkcija
	function sendRequest() {
		updateInputDataErrors();

		if (areAnyInputDataErrors()) {
			return;
		}

		sendCreateRequest();
	}

	function cancelCreate() {
		wantsToAddCoffee.set(false);
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	onclick={cancelCreate}
	id="overlay"
	class="fixed left-0 top-0 h-full w-full bg-black opacity-70"
></div>
<div
	id="modal"
	class="fixed left-1/2 top-1/2 h-fit w-96 -translate-x-1/2 -translate-y-1/2 transform
	rounded-lg bg-white px-4 pb-3 pt-2"
>
	<div>
		<p class="mb-5 mt-1 text-2xl font-medium">Naujos kavos pridėjimas</p>

		<div class="mb-8 flex flex-col gap-4">
			<div class="flex flex-col">
				<label
					for="name"
					class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
					>Pavadinimas</label
				>
				<input
					id="name"
					bind:value={name.value}
					type="text"
					class="{name.error
						? 'border-red-600'
						: 'border-gray-400'} rounded-md hover:border-gray-500 focus:border-black focus:outline-none focus:ring-0"
				/>
				{#if name.error}
					<p class="ml-1 mt-1 text-xs font-medium text-red-600">
						{name.errorMessage}
					</p>
				{/if}
			</div>

			<div class="flex flex-col">
				<label
					for="desc"
					class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
					>Aprašas</label
				>
				<input
					id="desc"
					bind:value={description.value}
					type="text"
					class="{description.error
						? 'border-red-600'
						: 'border-gray-400'} rounded-md hover:border-gray-500 focus:border-black focus:outline-none focus:ring-0"
				/>
				{#if description.error}
					<p class="ml-1 mt-1 text-xs font-medium text-red-600">
						{description.errorMessage}
					</p>
				{/if}
			</div>

			<div class="flex flex-col">
				<label
					for="roast_type"
					class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
					>Skrudinimo tipas</label
				>
				<select
					id="roast_type"
					bind:value={typeValue}
					class="rounded-md border-gray-400 hover:border-gray-500 focus:border-black focus:outline-none focus:ring-0"
				>
					<option value="light">Šviesus</option>
					<option value="medium">Vidutiniškas</option>
					<option value="dark">Tamsus</option>
				</select>
			</div>

			<div class="flex flex-col">
				<label
					for="stock"
					class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
					>Kiekis</label
				>
				<input
					id="stock"
					bind:value={stock.value}
					type="number"
					class="{stock.error
						? 'border-red-600'
						: 'border-gray-400'} rounded-md hover:border-gray-500 focus:border-black focus:outline-none focus:ring-0"
				/>
				{#if stock.error}
					<p class="ml-1 mt-1 text-xs font-medium text-red-600">
						{stock.errorMessage}
					</p>
				{/if}
			</div>

			<div class="flex flex-col">
				<label
					for="price"
					class="relative mb-1 ml-1 w-fit after:absolute after:-right-2.5 after:inline-block after:font-medium after:text-red-600 after:content-['*']"
					>Kaina</label
				>
				<input
					id="price"
					bind:value={price.value}
					type="number"
					class="{price.error
						? 'border-red-600'
						: 'border-gray-400'} rounded-md hover:border-gray-500 focus:border-black focus:outline-none focus:ring-0"
				/>
				{#if price.error}
					<p class="ml-1 mt-1 text-xs font-medium text-red-600">
						{price.errorMessage}
					</p>
				{/if}
			</div>
		</div>

		<div class="flex justify-end gap-3">
			<button
				onclick={cancelCreate}
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
</div>
