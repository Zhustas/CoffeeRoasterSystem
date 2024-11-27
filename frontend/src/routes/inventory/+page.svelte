<script lang="ts">
	import type { PageData } from './$types';
	import type { Coffee } from '$lib/Coffee';
	import NewCoffee from './NewCoffee.svelte';
	import { wantsToAddCoffee } from '../../stores';

	let { data }: { data: PageData } = $props();
	let coffees: Coffee[] = data.coffees;

	// ADD PROCESAS

	let wantsToAdd = $state();
	const unsubscribe = wantsToAddCoffee.subscribe((value) => {
		wantsToAdd = value;
	});

	function startAddProcess() {
		wantsToAddCoffee.set(true);
	}

	// DELETE PROCESAS

	let wantsToDelete = $state(false);
	let selectedID: number;

	function startDeleteProcess(ID: number) {
		wantsToDelete = true;
		selectedID = ID;

		// document.body.style.overflow = 'hidden';
	}

	function endDeleteProcess() {
		wantsToDelete = false;
		document.body.style.overflow = 'visible';
	}

	async function deleteCoffee(ID: number) {
		const response = await fetch(`/api/deletecoffee`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ ID: ID })
		});

		console.log(await response.json());
		endDeleteProcess();
	}
</script>

<article class="mx-64 my-10 flex flex-col">
	<!-- Parodyti confirmation box -->
	{#if wantsToDelete}
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<div
			onclick={endDeleteProcess}
			id="overlay"
			class="fixed left-0 top-0 h-full w-full bg-black opacity-60"
		></div>
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
			class="text-md bg-button grid h-12 grid-cols-8 items-center rounded-t-md border border-gray-300 text-center font-medium"
		>
			<p class="h-full content-center border-r border-gray-300">Pavadinimas</p>
			<p class="h-full content-center border-r border-gray-300">Aprašas</p>
			<p class="h-full content-center border-r border-gray-300">Skrudinimo tipas</p>
			<p class="h-full content-center border-r border-gray-300">Kiekis (g.)</p>
			<p class="h-full content-center border-r border-gray-300">Kaina (per 100 g.)</p>
			<p class="h-full content-center border-r border-gray-300">Sukurta</p>
			<p class="h-full content-center border-r border-gray-300">Atnaujinta</p>
			<p class="h-full content-center border-gray-300"></p>
		</header>
		<div>
			{#if coffees.length !== 0}
				{#each coffees as coffee, i}
					<div
						class="grid h-fit grid-cols-8 items-center border-b border-l border-r border-gray-300 text-center hover:bg-slate-600 hover:text-white {i +
							1 ===
						coffees.length
							? 'rounded-b-md'
							: ''} {i % 2 === 0 ? 'bg-main' : 'bg-main'}"
					>
						<p class="h-full content-center break-words border-r border-gray-300 px-2 py-1">
							{coffee.name}
						</p>
						<p class="h-full content-center break-words border-r border-gray-300 px-2 py-1">
							{coffee.description}
						</p>
						<p class="h-full content-center break-words border-r border-gray-300 px-2 py-1">
							{coffee.roast_type}
						</p>
						<p class="h-full content-center break-words border-r border-gray-300 px-2 py-1">
							{coffee.stock}
						</p>
						<p class="h-full content-center break-words border-r border-gray-300 px-2 py-1">
							{coffee.price}
						</p>
						<p class="h-full content-center break-words border-r border-gray-300 px-2 py-1">-</p>
						<p class="h-full content-center break-words border-r border-gray-300 px-2 py-1">-</p>
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
					class="grid h-10 grid-cols-1 items-center rounded-b-md border-b border-l border-r border-gray-300 bg-white text-center"
				>
					<p>Kavos nėra</p>
				</div>{/if}
		</div>
	</div>
</article>
