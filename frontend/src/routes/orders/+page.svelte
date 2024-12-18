<script lang="ts">
	import type { PageData } from './$types';
	import type { Order } from '$lib/order';
	import Footer from '../../components/Footer.svelte';
	import { onMount } from 'svelte';
	import Header from '$lib/components/Header.svelte';
	import { USER_ADMIN, USER_ROASTER } from '$lib/constans';
	import AlertMessage from '../../lib/components/AlertMessage.svelte';
	import * as AlertMessageConstants from '../../constants/AlertMessageConstants';

	let { data }: { data: PageData } = $props();
	let orders: Order[] = $state(data.orders.orders);
	let user = data.user;

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

	let wantsToChangeStatus = $state(false);
	let orderIdChangeStatus: number;
	let orderStatusChangeStatus: string;

	function startChangeStatusProcess(id: number, status: string) {
		wantsToChangeStatus = true;
		orderIdChangeStatus = id;
		orderStatusChangeStatus = status;
	}

	function endChangeStatusProcess() {
		wantsToChangeStatus = false;
	}

	function endDeleteProcess() {
		wantsToChangeStatus = false;
	}

	async function sendRequest() {
		const response = await fetch(`/api/changeorderstatus`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ ID: orderIdChangeStatus, status: orderStatusChangeStatus })
		});

		if (response.ok) {
			showAlertMessage(
				AlertMessageConstants.STATUS_SUCCESS,
				`Užsakymo Nr. ${orderIdChangeStatus} statusas pakeistas į "${STATUS_LT[orderStatusChangeStatus]}".`
			);
		} else {
			showAlertMessage(
				AlertMessageConstants.STATUS_SUCCESS,
				`Įvyko klaidą keičiant užsakymo Nr. ${orderIdChangeStatus} statusą.`
			);
		}

		endChangeStatusProcess();
		getOrders();
	}

	async function getOrders() {
		const response = await fetch(`/api/getallorders`, {
			method: 'GET'
		});

		const body = await response.json();
		orders = body.orders.orders;
		filterOrders();
	}

	async function sendDeleteRequest() {
		const response = await fetch(`/api/deleteorder/${orderIdChangeStatus}`, {
			method: 'POST'
		});

		if (response.ok) {
			showAlertMessage(AlertMessageConstants.STATUS_SUCCESS, 'Užsakymas ištrintas.');
		} else {
			showAlertMessage(AlertMessageConstants.STATUS_FAILURE, 'Įvyko klaida ištrinant užsakymą.');
		}

		endDeleteProcess();
		getOrders();
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
</script>

<svelte:head>
	<title>Visi užsakymai - Kavos Skrudykla</title>
</svelte:head>

{#if wantsToChangeStatus}
	<button
		aria-label="Paspauskite, kad baigtumėte ištrynimo procesą."
		onclick={endChangeStatusProcess}
		id="overlay"
		class="fixed left-0 top-0 z-10 h-full w-full cursor-default bg-black opacity-60"
	></button>
	<div
		id="modal"
		class="fixed left-1/2 top-1/2 z-10 h-fit w-[32rem] -translate-x-1/2 -translate-y-1/2 transform rounded-lg bg-white px-4 pb-3 pt-2"
	>
		<div>
			<p class="mb-2 mt-1 text-2xl font-medium">Užsakymo ištrynimas/statuso pakeitimas</p>
			<p class="mb-2">Keičiate užsakymo Nr. {orderIdChangeStatus} statusą:</p>
			<select
				bind:value={orderStatusChangeStatus}
				class="mb-7 rounded-md hover:border-gray-400 focus:border-black focus:outline-none focus:ring-0"
			>
				<option value="pending">Naujas</option>
				<option value="in_progress">Vykdomas</option>
				<option value="shipped">Išsiųstas</option>
				<option value="completed">Įvykdytas</option>
				<option value="cancelled">Atšauktas</option>
			</select>
			<div class="flex">
				<button
					onclick={sendDeleteRequest}
					class="rounded bg-red-500 px-6 py-1 font-medium text-white hover:bg-red-600"
					>Ištrinti</button
				>
				<div class="flex grow justify-end gap-2">
					<button
						onclick={endChangeStatusProcess}
						class="10 rounded border border-gray-600 bg-white px-6 py-1 font-medium text-gray-600 hover:bg-gray-200"
						>Atšaukti</button
					>
					<button
						onclick={sendRequest}
						class="rounded bg-green-500 px-6 py-1 font-medium text-white hover:bg-green-600"
						>Pakeisti</button
					>
				</div>
			</div>
		</div>
	</div>
{/if}

<div class="flex min-h-screen flex-col">
	{#if user.role === USER_ROASTER}
		<Header userType={USER_ROASTER} />
	{:else if user.role === USER_ADMIN}
		<Header userType={USER_ADMIN} />
	{/if}

	<article class="mx-page mt-top mb-bottom grow">
		<div class="mb-8 border-b-2 border-black">
			<p class="mb-2 ml-10 text-4xl font-medium">Visi užsakymai</p>
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

					<div id="statuses" class="grow">
						<p class="mb-1 ml-1">Statusas</p>
						<div class="flex">
							{#each Object.values(STATUS_LT) as value, i}
								<div class="mr-5 flex h-10">
									<input
										id={value.toLocaleLowerCase()}
										type="checkbox"
										bind:checked={fieldCheckBoxes[i]}
										class="size-5 self-center rounded-md bg-gray-200 focus:outline-none focus:ring-0"
									/>
									<label for={value.toLocaleLowerCase()} class="ml-1.5 self-center">{value}</label>
								</div>
							{/each}
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
			<div class="flex flex-col">
				{#if ordersToShow.length !== 0}
					{#each ordersToShow as order, i}
						<button
							onclick={() => startChangeStatusProcess(order.id, order.status)}
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
						</button>
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

{#if alertMessageHandler.show}
	<AlertMessage status={alertMessageHandler.status} message={alertMessageHandler.message} />
{/if}
