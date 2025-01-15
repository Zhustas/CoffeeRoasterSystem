<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import type { Order } from '$lib/order';
	import Footer from '$lib/components/Footer.svelte';
	import { wantsToCancelOrder } from '$lib/stores';
	import AlertMessage from '$lib/components/AlertMessage.svelte';
	import * as AlertMessageConstants from '$lib/constants/AlertMessageConstants';
	import Header from '$lib/components/Header.svelte';
	import { USER_CUSTOMER } from '$lib/constants/UserTypeConstants';
	import type { Coffee } from '$lib/coffee';

	let { data }: { data: PageData } = $props();
	const orders: Order[] = $state(data.orders);
	const orderItems = data.orderItems;
	const coffees: Coffee[] = data.coffees;

	const dates: string[] = $state([]);

	const STATUS_LT = {
		pending: 'Užsakytas',
		in_progress: 'Vykdomas',
		shipped: 'Išsiųstas',
		completed: 'Įvykdytas',
		cancelled: 'Atšauktas'
	};

	let euroFormatter: Intl.NumberFormat = $state();

	onMount(() => {
		filterDates();

		euroFormatter = new Intl.NumberFormat(navigator.language, {
			style: 'currency',
			currency: 'EUR'
		});
	});

	function formatDate(date: Date) {
		const newDate = new Date(date);

		const month = `${newDate.getMonth() + 1}`.padStart(2, '0');
		const day = `${newDate.getDate()}`.padStart(2, '0');

		return `${newDate.getFullYear()}-${month}-${day}`;
	}

	function filterDates() {
		orders.forEach((order) => {
			const formattedDate = formatDate(order.created_at);

			if (!dates.includes(formattedDate)) {
				dates.push(formattedDate);
			}
		});

		dates.reverse();
	}

	let wantsToCancel = $state();
	const unsubscribe = wantsToCancelOrder.subscribe((value) => {
		wantsToCancel = value;
	});

	function canCancelOrder(status: string) {
		if (STATUS_LT[status] === 'Užsakytas') {
			return true;
		}

		return false;
	}

	let cancelID: number = -1;

	function startCancelingOrder(id: number) {
		wantsToCancelOrder.set(true);
		cancelID = id;
	}

	function cancelCancelingOrder() {
		wantsToCancelOrder.set(false);
		cancelID = -1;
	}

	async function sendRequest() {
		const response = await fetch(`/api/cancelorder`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ ID: cancelID })
		});

		if (response.ok) {
			const index = orders.findIndex((order) => order.id === cancelID);
			orders[index].status = 'cancelled';

			cancelCancelingOrder();
			showAlertMessage(AlertMessageConstants.STATUS_SUCCESS, 'Užsakymas sėkmingai atšauktas.');
		}
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

	function getCoffeeNameByOrderId(orderId: number) {
		const orderItem = orderItems.find((orderItem) => orderItem.order_id === orderId);
		const coffee = coffees.find((coffee) => orderItem && coffee.ID === orderItem.coffee_id);

		return coffee ? coffee.name : 'Įvyko klaida';
	}

	function getQuantityByOrderId(orderId: number) {
		const orderItem = orderItems.find((orderItem) => orderItem.order_id === orderId);

		return orderItem ? orderItem.quantity : '0';
	}

	function convertToRoastTypeLT(roastType: string) {
		if (roastType === 'light') return 'Šviesus';
		if (roastType === 'medium') return 'Vidutinis';
		if (roastType === 'dark') return 'Tamsus';
	}

	function getCoffeeRoastTypeByOrderId(orderId: number) {
		const orderItem = orderItems.find((orderItem) => orderItem.order_id === orderId);
		const coffee = coffees.find((coffee) => orderItem && coffee.ID === orderItem.coffee_id);

		return coffee ? coffee.roast_type : 'Įvyko klaida';
	}

	function getCoffeeRoastTypeByOrderIdAndConvertToLT(orderId: number) {
		const orderItem = orderItems.find((orderItem) => orderItem.order_id === orderId);
		const coffee = coffees.find((coffee) => orderItem && coffee.ID === orderItem.coffee_id);

		return coffee ? convertToRoastTypeLT(coffee.roast_type) : 'Įvyko klaida';
	}

	function getPrice(orderId: number) {
		// coffee.price * (quantity / 100)
		const orderItem = orderItems.find((orderItem) => orderItem.order_id === orderId);
		const coffee = coffees.find((coffee) => orderItem && coffee.ID === orderItem.coffee_id);

		return coffee.price * (orderItem.quantity / 100);
	}
</script>

<svelte:head>
	<title>Mano užsakymai - Kavos Skrudykla</title>
</svelte:head>

{#if wantsToCancel}
	<button
		onclick={cancelCancelingOrder}
		id="overlay"
		aria-label="Cancel canceling order process"
		class="fixed left-0 top-0 h-full w-full cursor-default bg-black opacity-70"
	></button>
	<div
		id="modal"
		class="fixed left-1/2 top-1/2 h-fit w-96 -translate-x-1/2 -translate-y-1/2 transform
	rounded-lg bg-white px-6 pb-3 pt-2"
	>
		<p class="mb-4 mt-1 text-2xl font-medium">Užsakymo atšaukimas</p>
		<p class="mb-5">Ar tikrai norite atšaukti užsakymą?</p>
		<div class="flex justify-end gap-3">
			<button
				onclick={cancelCancelingOrder}
				class="10 rounded border border-gray-600 bg-white px-6 py-1 font-medium text-gray-600 hover:bg-gray-200"
				>Atgal</button
			>
			<button
				onclick={sendRequest}
				class="rounded bg-red-500 px-6 py-1 font-medium text-white hover:bg-red-600"
				>Atšaukti</button
			>
		</div>
	</div>
{/if}

<div class="flex min-h-screen flex-col">
	<div class="z-10">
		<Header userType={USER_CUSTOMER} />
	</div>

	<article class="mx-page mb-bottom mt-top grow">
		<div class="mb-8 border-b-2 border-black">
			<p class="mb-2 ml-10 text-4xl font-medium">Mano užsakymai</p>
		</div>

		<div class="px-10">
			{#if orders.length !== 0}
				{#each dates as date, i}
					<div class={i !== dates.length - 1 ? 'mb-8' : ''}>
						<p class="mb-4 ml-2 w-96 border-b-2 border-black text-2xl font-medium">{date}</p>
						<div class="flex flex-col gap-8">
							{#each orders as order}
								{#if formatDate(order.created_at) === date}
									<div class="flex min-h-32 rounded-md bg-white py-2 hover:bg-gray-100">
										<div class="flex grow">
											{#if getCoffeeRoastTypeByOrderId(order.id) === 'light'}
												<img alt="Light coffee bean" src="light_bean.png" class="w-36 p-2" />
											{:else if getCoffeeRoastTypeByOrderId(order.id) === 'medium'}
												<img alt="Medium coffee bean" src="medium_bean.png" class="w-36 p-2" />
											{:else if getCoffeeRoastTypeByOrderId(order.id) === 'dark'}
												<img alt="Medium coffee bean" src="dark_bean.png" class="w-36 p-2" />
											{/if}

											<div>
												<p class="mt-8 w-[28rem] px-2 text-xl font-medium">
													{getCoffeeNameByOrderId(order.id)}
												</p>
											</div>
										</div>

										<div class="flex justify-end gap-5">
											<div class=" flex h-2/3 flex-col justify-center self-center border-r-2">
												<div class="mr-5">
													<p class="text-xs text-gray-500">Skrudinimo tipas</p>
													<p class="font-medium">
														{getCoffeeRoastTypeByOrderIdAndConvertToLT(order.id)}
													</p>
												</div>
											</div>

											<div class=" flex h-2/3 flex-col justify-center self-center border-r-2">
												<div class="mr-5">
													<p class="text-xs text-gray-500">Kiekis</p>
													<p class="font-medium">{getQuantityByOrderId(order.id)} g.</p>
												</div>
											</div>

											<div class="flex h-2/3 flex-col justify-center self-center border-r-2">
												<p class="text-xs text-gray-500">Kaina</p>
												<p class="mr-5 font-medium">
													{euroFormatter.format(getPrice(order.id))}
												</p>
											</div>

											<div class=" flex h-2/3 flex-col justify-center self-center border-r-2">
												<div class="mr-5">
													<p class="text-xs text-gray-500">Užsakymo Nr.</p>
													<p class="font-medium">{order.id}</p>
												</div>
											</div>

											<div class="relative mr-5 flex h-2/3 flex-col justify-center self-center">
												<p class="text-xs text-gray-500">Statusas</p>
												<p class="font-medium">{STATUS_LT[order.status]}</p>

												{#if canCancelOrder(order.status)}
													<button
														onclick={() => startCancelingOrder(order.id)}
														class="absolute left-4 top-24 text-xs underline">Atšaukti</button
													>
												{/if}
											</div>
										</div>
									</div>
								{/if}
							{/each}
						</div>
					</div>
				{/each}
			{:else}
				<div class="flex justify-center">
					<p class="font-medium">Jūs dar nepateikėte nei vieno užsakymo.</p>
				</div>
			{/if}
		</div>
	</article>

	<Footer />
</div>

{#if alertMessageHandler.show}
	<AlertMessage status={alertMessageHandler.status} message={alertMessageHandler.message} />
{/if}
