<script lang="ts">
	import type { PageData } from './$types';
	import type { Coffee } from '$lib/coffee';
	import NewCoffee from '$lib/components/NewCoffee.svelte';
	import { wantsToAddCoffee } from '$lib/stores';
	import Footer from '$lib/components/Footer.svelte';
	import Header from '$lib/components/Header.svelte';
	import { USER_ROASTER } from '$lib/constants/UserTypeConstants';
	import { onMount } from 'svelte';
	import AlertMessage from '$lib/components/AlertMessage.svelte';
	import * as AlertMessageConstants from '$lib/constants/AlertMessageConstants';
	import { alertStore, type AlertMessageHandler, showAlertMessage } from '$lib/alertmessagehandler';

	let alertMessageHandler: AlertMessageHandler = $state();
	alertStore.subscribe((state) => {
		alertMessageHandler = state;
	});

	let { data }: { data: PageData } = $props();
	let coffees: Coffee[] = data.coffees;

	let euroFormatter: Intl.NumberFormat = $state();

	onMount(() => {
		euroFormatter = new Intl.NumberFormat(navigator.language, {
			style: 'currency',
			currency: 'EUR'
		});
	});

	// Add procesas
	let wantsToAdd = $state();
	const unsubscribe = wantsToAddCoffee.subscribe((value) => {
		wantsToAdd = value;
	});

	function startAddProcess() {
		wantsToAddCoffee.set(true);
	}

	// Delete procesas
	let wantsToDelete = $state(false);
	let selectedID: number;

	function startDeleteProcess(ID: number) {
		wantsToDelete = true;
		selectedID = ID;
	}

	function endDeleteProcess() {
		wantsToDelete = false;
		document.body.style.overflow = 'visible';
	}

	async function deleteCoffee(ID: number) {
		const response = await fetch(`/api/deletecoffee`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ ID })
		});

		if (response.ok) {
			showAlertMessage(
				AlertMessageConstants.STATUS_SUCCESS,
				`Kava (ID - ${ID}) sėkmingai ištrinta.`
			);
		} else {
			showAlertMessage(
				AlertMessageConstants.STATUS_FAILURE,
				`Kavos (ID - ${ID}) nepavyko ištrinti.`
			);
		}
		endDeleteProcess();
	}

	// Stulpeliai
	const numberOfColumns = 8;
	const namesOfColumns: string[] = [
		'Pavadinimas',
		'Aprašas',
		'Skrudinimo tipas',
		'Kiekis (g.)',
		'Kaina',
		'Sukurta',
		'Atnaujinta'
	];
</script>

<svelte:head>
	<title>Kavos inventorius - Kavos Skrudykla</title>
</svelte:head>

<div class="flex min-h-screen flex-col">
	<Header userType={USER_ROASTER} />

	<article class="mx-page mb-bottom mt-top flex grow flex-col">
		<div class="mb-8 border-b-2 border-black">
			<p class="mb-2 ml-10 text-4xl font-medium">Kavos inventorius</p>
		</div>

		{#if wantsToDelete}
			<button
				aria-label="Paspauskite, kad baigtumėte ištrynimo procesą."
				onclick={endDeleteProcess}
				id="overlay"
				class="fixed left-0 top-0 h-full w-full cursor-default bg-black opacity-60"
			></button>
			<div
				id="modal"
				class="fixed left-1/2 top-1/2 h-fit w-96 -translate-x-1/2 -translate-y-1/2 transform rounded-lg bg-white px-4 pb-3 pt-2"
			>
				<div>
					<p class="mb-2 mt-1 text-2xl font-medium">Ištrynimas</p>
					<p class="mb-7">Ar tikrai norite ištrinti pasirinktą kavą?</p>
					<div class="flex justify-end gap-3">
						<button
							onclick={endDeleteProcess}
							class="10 rounded border border-gray-600 bg-white px-6 py-1 font-medium text-gray-600 hover:bg-gray-200"
							>Ne</button
						>
						<button
							onclick={() => deleteCoffee(selectedID)}
							class="rounded bg-red-500 px-6 py-1 font-medium text-white hover:bg-red-600"
							>Taip</button
						>
					</div>
				</div>
			</div>
		{/if}

		{#if wantsToAdd}
			<NewCoffee />
		{/if}

		<button
			onclick={startAddProcess}
			class="mb-5 mr-2 w-48 self-end rounded-md bg-green-500 py-1 font-medium text-white hover:bg-green-600"
		>
			Sukurti naują kavą</button
		>

		<div class="flex flex-col">
			<header
				class="text-md mb-2 grid h-12 grid-cols-8 items-center rounded-md border border-gray-300 bg-white text-center font-medium"
			>
				{#each Array(numberOfColumns) as _, i}
					{#if i !== numberOfColumns - 1}
						<p class="h-full content-center border-r border-gray-300">{namesOfColumns[i]}</p>
					{:else}
						<p class="h-full content-center border-gray-300">{namesOfColumns[i]}</p>
					{/if}
				{/each}
			</header>
			<div>
				{#if coffees.length !== 0}
					{#each coffees as coffee, i}
						<div
							class="grid h-fit grid-cols-8 items-center border-b border-l border-r border-black text-center hover:bg-gray-400 hover:text-white {i +
								1 ===
							coffees.length
								? 'rounded-b-md'
								: ''} {i % 2 === 0 ? 'bg-gray-200' : 'bg-gray-300'} {i === 0
								? 'rounded-t-md border-t'
								: ''}"
						>
							<p class="h-full content-center break-words border-r border-black px-2 py-1">
								{coffee.name}
							</p>
							<p class="h-full grow content-center break-words border-r border-black px-2 py-1">
								{coffee.roast_type}
							</p>
							<p class="h-full content-center break-words border-r border-black px-2 py-1">
								{coffee.description}
							</p>
							<p class="h-full content-center break-words border-r border-black px-2 py-1">
								{coffee.stock}
							</p>
							<p class="h-full content-center break-words border-r border-black px-2 py-1">
								{#if euroFormatter}
									{euroFormatter.format(coffee.price)}
								{:else}
									Formatuojama...
								{/if}
							</p>
							<p class="h-full content-center break-words border-r border-black px-2 py-1">-</p>
							<p class="h-full content-center break-words border-r border-black px-2 py-1">-</p>
							<div
								class="flex h-full content-center items-center justify-center border-gray-300 px-2 py-1"
							>
								<button
									onclick={() => startDeleteProcess(coffee.ID)}
									class="h-8 w-16 rounded-md bg-red-500 text-center text-xs font-medium text-white hover:bg-red-600"
								>
									Ištrinti
								</button>
							</div>
						</div>
					{/each}
				{:else}
					<div
						class="grid h-10 grid-cols-1 items-center rounded-md border border-gray-300 bg-white text-center"
					>
						<p>Kavos nėra</p>
					</div>{/if}
			</div>
		</div>
	</article>

	<Footer />
</div>

{#if alertMessageHandler.show}
	<AlertMessage status={alertMessageHandler.status} message={alertMessageHandler.message} />
{/if}
