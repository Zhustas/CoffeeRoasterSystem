<script lang="ts">
	import type { PageData } from './$types';
	import type { Order } from '$lib/order';
	import Footer from '../../components/Footer.svelte';
	import { onMount } from 'svelte';
	import Header from '$lib/components/Header.svelte';
	import { USER_ROASTER } from '$lib/constans';

	let { data }: { data: PageData } = $props();
	let orders: Order[] = data.orders.orders;

	let euroFormatter: Intl.NumberFormat = $state();

	onMount(() => {
		euroFormatter = new Intl.NumberFormat(navigator.language, {
			style: 'currency',
			currency: 'EUR'
		});

		ordersToShow = orders;
	});

	const numberOfColumns = 5;
	const namesOfColumns: string[] = ['Užsakymo ID', 'Statusas', 'Kaina', 'Pirkėjo ID', 'Sukurtas'];

	const STATUS_LT = {
		pending: 'Naujas',
		in_progress: 'Vykdomas',
		shipped: 'Išsiųstas',
		completed: 'Įvykdytas',
		cancelled: 'Atšauktas'
	};

	function formatDate(date: Date) {
		const newDate = new Date(date);

		const hour = `${newDate.getHours()}`.padStart(2, '0');
		const minute = `${newDate.getMinutes()}`.padStart(2, '0');
		const seconds = `${newDate.getSeconds()}`.padStart(2, '0');

		const formattedDate = `${newDate.getFullYear()}-${newDate.getMonth() + 1}-${newDate.getDate()} ${hour}:${minute}:${seconds}`;

		return formattedDate;
	}

	// Laukai
	let fieldID: number | null = $state();
	let fieldUserID: number | null = $state();
	let fieldCheckBoxes: boolean[] = $state([false, false, false, false, false]);
	let fieldDateFrom: Date | null = $state();
	let fieldDateUntil: Date | null = $state();

	let ordersToShow: Order[] = $state([]);

	const STATUSES: string[] = ['pending', 'in_progress', 'shipped', 'completed', 'cancelled'];

	function filterOrders() {
		ordersToShow = orders;

		if (fieldID) {
			ordersToShow = ordersToShow.filter((order) => order.id === fieldID);
		}

		if (fieldUserID) {
			ordersToShow = ordersToShow.filter((order) => order.user_id === fieldUserID);
		}

		if (fieldCheckBoxes.some((value) => value)) {
			const activeStatuses = fieldCheckBoxes
				.map((value, i) => (value ? STATUSES[i] : null))
				.filter((status) => status !== null);

			ordersToShow = ordersToShow.filter((order) => activeStatuses.includes(order.status));
		}

		if (fieldDateFrom) {
			ordersToShow = ordersToShow.filter(
				(order) => new Date(fieldDateFrom).getTime() <= new Date(order.created_at).getTime()
			);
		}

		if (fieldDateUntil) {
			ordersToShow = ordersToShow.filter(
				(order) => new Date(order.created_at).getTime() <= new Date(fieldDateUntil).getTime()
			);
		}
	}

	function clearFilters() {
		fieldID = null;
		fieldUserID = null;
		for (let i = 0; i < fieldCheckBoxes.length; i++) {
			fieldCheckBoxes[i] = false;
		}
		fieldDateFrom = null;
		fieldDateUntil = null;

		filterOrders();
	}
</script>

<svelte:head>
	<title>Nauji užsakymai - Kavos Skrudykla</title>
</svelte:head>

<div class="flex min-h-screen flex-col">
	<Header userType={USER_ROASTER} />

	<article class="mx-page mt-top mb-bottom grow">
		<div class="mb-8 border-b-2 border-black">
			<p class="mb-2 ml-10 text-4xl font-medium">Nauji užsakymai</p>
		</div>

		<div class="mb-7 flex flex-col">
			<div id="filters" class="mb-5 flex gap-16">
				<div id="ids" class="flex flex-col gap-4">
					<div>
						<p class="mb-1 ml-1">Užsakymo ID</p>
						<input
							type="number"
							bind:value={fieldID}
							placeholder="Įveskite užsakymo ID"
							class="rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
						/>
					</div>
					<div>
						<p class="mb-1 ml-1">Pirkėjo ID</p>
						<input
							type="number"
							bind:value={fieldUserID}
							placeholder="Įveskite pirkėjo ID"
							class="rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
						/>
					</div>
				</div>

				<div id="other" class="flex grow flex-col gap-4">
					<div id="dates">
						<p class="mb-1 ml-1">Sukurtas</p>
						<div class="flex gap-5">
							<div>
								<label for="from" class="mr-2">Nuo</label>
								<input
									id="from"
									type="datetime-local"
									bind:value={fieldDateFrom}
									class="rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
								/>
							</div>
							<div>
								<label for="until" class="mr-2">Iki</label>
								<input
									id="until"
									type="datetime-local"
									bind:value={fieldDateUntil}
									class="rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
								/>
							</div>
						</div>
					</div>
				</div>
			</div>

			<div id="buttons" class="flex justify-between">
				<button
					onclick={clearFilters}
					class="ml-2 rounded-md bg-gray-200 px-4 py-1 font-medium hover:bg-slate-300"
					>Išvalyti filtrus</button
				>
				<button
					onclick={filterOrders}
					class="bg-button hover:bg-hover-button mr-2 rounded-md px-4 py-1 font-medium"
					>Filtruoti</button
				>
			</div>
		</div>

		<div class="flex flex-col">
			<header
				class="text-md mb-2 grid h-12 grid-cols-5 items-center rounded-md border border-gray-300 bg-white text-center font-medium"
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
				{#if ordersToShow.length !== 0}
					{#each ordersToShow as order, i}
						<div
							class="grid h-fit grid-cols-5 items-center border-b border-l border-r border-black text-center hover:bg-gray-400 hover:text-white {i +
								1 ===
							orders.length
								? 'rounded-b-md'
								: ''} {i % 2 === 0 ? 'bg-gray-200' : 'bg-gray-300'} {i === 0
								? 'rounded-t-md border-t'
								: ''}"
						>
							<p class="h-full content-center break-words border-r border-black px-2 py-1">
								{order.id}
							</p>
							<p class="h-full content-center break-words border-r border-black px-2 py-1">
								{STATUS_LT[order.status]}
							</p>
							<p class="h-full content-center break-words border-r border-black px-2 py-1">
								{#if euroFormatter}
									{euroFormatter.format(order.total_amount)}
								{:else}
									Formatuojama...
								{/if}
							</p>
							<p class="h-full content-center break-words border-r border-black px-2 py-1">
								{order.user_id}
							</p>
							<p class="h-full content-center break-words border-black px-2 py-1">
								{formatDate(order.created_at)}
							</p>
						</div>
					{/each}
				{:else}
					<div
						class="grid h-10 grid-cols-1 items-center rounded-md border border-gray-300 bg-white text-center"
					>
						<p class="text-sm font-medium">Užsakymų nėra</p>
					</div>
				{/if}
			</div>
		</div>
	</article>

	<Footer />
</div>
