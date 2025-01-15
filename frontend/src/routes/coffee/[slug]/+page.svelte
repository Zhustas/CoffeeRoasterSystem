<script lang="ts">
	import { onMount } from 'svelte';
	import type { Coffee } from '$lib/coffee';
	import type { PageData } from './$types';
	import AlertMessage from '$lib/components/AlertMessage.svelte';
	import * as AlertMessageConstants from '$lib/constants/AlertMessageConstants';
	import CardCredentialsTab from '$lib/components/CardCredentialsTab.svelte';
	import { wantsToMakeOrder } from '$lib/stores';

	let { data }: { data: PageData } = $props();
	let coffee: Coffee = $state(data.coffee);

	let quantity: number = $state();

	let euroFormatter: Intl.NumberFormat = $state();

	onMount(() => {
		euroFormatter = new Intl.NumberFormat(navigator.language, {
			style: 'currency',
			currency: 'EUR'
		});
	});

	function changeToLT(roast_type: string) {
		if (roast_type === 'light') return 'šviesiai skrudinta';
		if (roast_type === 'medium') return 'vidutiniškai skrudinta';
		if (roast_type === 'dark') return 'tamsiai skrudinta';
	}

	async function makeOrder() {
		coffee.stock -= quantity;

		const response = await fetch(`/api/neworder`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ items: [{ coffee_id: coffee.ID, quantity: Number(quantity) }] })
		});

		if (response.ok) {
			showAlertMessage(AlertMessageConstants.STATUS_SUCCESS, 'Užsakymas pateiktas!');
			// Gauti kavos objektą iš DB
			getCoffee();
		} else {
			showAlertMessage(AlertMessageConstants.STATUS_FAILURE, 'Įvyko klaida pateikiant užsakymą.');
		}

		quantity = AVAILABLE_SELLING_OPTIONS[0];
	}

	async function getCoffee() {
		const response = await fetch(`/api/getcoffee/${coffee.ID}`, {
			method: 'POST'
		});

		const body = await response.json();
		coffee = body.coffee;
	}

	interface AlertMessageHandler {
		show: boolean;
		status: string;
		message: string;
	}

	const alertMessageHandler: AlertMessageHandler = $state({ show: false, status: '', message: '' });
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

	const AVAILABLE_SELLING_OPTIONS: number[] = [100, 200, 500, 1000];

	let wantsToOrder = $state();
	const unsubscribe = wantsToMakeOrder.subscribe((value) => {
		wantsToOrder = value;
	});

	function startOrdering() {
		wantsToMakeOrder.set(true);
	}
</script>

<svelte:head>
	<title>{coffee.name} - Kavos Skrudykla</title>
</svelte:head>

{#if wantsToOrder}
	<CardCredentialsTab {makeOrder} />
{/if}

<div class="flex min-h-screen flex-col">
	<header>
		<nav class="flex h-16 border-b border-gray-400 bg-white px-64">
			<a aria-label="Perkrauti šį puslapį" href="/main" class="self-center p-2">
				<img alt="Tekstas - Kavos Skrudykla" src="../coffee_roaster_title.png" class="h-12" />
			</a>
			<div class="flex grow items-center justify-end font-medium">
				<a href="/subscription" class="mx-4 p-2">Prenumerata</a>
				<a href="/history" class="mx-4 p-2">Užsakymai</a>
				<a href="/account" class="mx-4 p-2">Paskyra</a>
				<div class="h-full w-px bg-gray-400"></div>
				<a href="/faq" class="mx-4 p-2">DUK</a>
			</div>
		</nav>
	</header>
	<article class="mx-64 flex grow">
		<div id="image" class="flex basis-1/2 justify-center px-5 pb-20 pt-10">
			<div class="rounded-md bg-white p-5">
				{#if coffee.roast_type === 'light'}
					<img alt="Light coffee bean" src="../../light_bean.png" class="h-full" />
				{:else if coffee.roast_type === 'medium'}
					<img alt="Medium roast coffee bean" src="../medium_bean.png" class="h-full" />
				{:else if coffee.roast_type === 'dark'}
					<img alt="Dark roast coffee bean" src="../dark_bean.png" class="h-full" />
				{/if}
			</div>
		</div>
		<div id="information" class="flex basis-1/2 px-5 pb-20 pt-10">
			<div class="flex grow flex-col rounded-md bg-white p-5">
				<p class="mb-5 text-4xl font-medium">{coffee.name}</p>
				<div class="mb-8 h-44 rounded-md bg-gray-100 px-2 py-1">
					<p class="text-sm">{coffee.description}</p>
				</div>
				<p class="mb-2">Tipas: <span class="font-medium">{changeToLT(coffee.roast_type)}</span></p>

				{#if coffee.stock >= AVAILABLE_SELLING_OPTIONS[0]}
					<div class="mb-8 flex">
						<label for="roast_type" class="mr-2 self-center">Pasirinkite kiekį:</label>
						<select
							bind:value={quantity}
							id="roast_type"
							class="mr-1 rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
						>
							{#each AVAILABLE_SELLING_OPTIONS as price}
								{#if coffee.stock >= price}
									<option value={price}>{price}</option>
								{:else}
									<option disabled value={price}>{price}</option>
								{/if}
							{/each}
						</select>
						<p class="self-center font-medium">g.</p>
					</div>
					<div class="flex grow justify-end">
						<div class="flex flex-col justify-end pb-5 text-end">
							<p class="text-sm text-gray-500">Kaina</p>
							<p class="text-3xl font-medium">
								{#if euroFormatter}
									{euroFormatter.format(Number(coffee.price * (quantity / 100)))}
								{:else}
									Kaina kraunama...
								{/if}
							</p>
						</div>
					</div>

					<div class="flex flex-col justify-end">
						<button
							onclick={startOrdering}
							class="h-20 w-full rounded-md bg-button text-2xl font-medium tracking-wider hover:bg-hover-button"
							>PIRKTI</button
						>
					</div>
				{:else}
					<div class="flex grow justify-center">
						<p class="self-center text-3xl font-medium text-red-500">Išparduota</p>
					</div>
				{/if}
			</div>
		</div>
	</article>
	<footer class="flex h-12 border-t border-gray-400 bg-white">
		<img
			alt="Kavos Skrudykla logo"
			src="../logo.png"
			class="ml-5 h-7 w-7 self-center rounded-full"
		/>
		<p class="ml-5 self-center text-xs font-medium text-gray-400">
			&copy; 2024 UAB "Kavos Skrudykla"
		</p>
	</footer>
</div>

{#if alertMessageHandler.show}
	<AlertMessage status={alertMessageHandler.status} message={alertMessageHandler.message} />
{/if}
