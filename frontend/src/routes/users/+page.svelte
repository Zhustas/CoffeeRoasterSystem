<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import AddUser from '$lib/components/AddUser.svelte';
	import { USER_ADMIN } from '$lib/constants/UserTypeConstants';
	import { userAddedStatus, wantsToAddUser } from '$lib/stores';
	import Footer from '$lib/components/Footer.svelte';
	import AlertMessage from '$lib/components/AlertMessage.svelte';
	import * as AlertMessageConstants from '$lib/constants/AlertMessageConstants';
	import type { PageData } from './$types';
	import { onMount } from 'svelte';

	let { data }: { data: PageData } = $props();
	let users = $state(data.users);

	let wantsToAdd: boolean = $state();
	wantsToAddUser.subscribe((value) => {
		wantsToAdd = value;
	});

	function startAddUserProcess() {
		wantsToAddUser.set(true);
	}

	interface AlertMessageHandler {
		show: boolean;
		status: string;
		message: string;
	}

	const alertMessageHandler: AlertMessageHandler = { show: false, status: '', message: '' };
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

	userAddedStatus.subscribe((value) => {
		if (value === 200) {
			showAlertMessage(AlertMessageConstants.STATUS_SUCCESS, 'Vartotojas pridėtas sėkmingai!');
		}
	});

	let wantsToDelete = $state(false);
	let selectedToDeleteId: number;
	let selectedToDeleteUsername: string;

	function startDeleteProcess(id: number, username: string) {
		wantsToDelete = true;
		selectedToDeleteId = id;
		selectedToDeleteUsername = username;
	}

	function cancelDeleteProcess() {
		wantsToDelete = false;
	}

	async function sendDeleteRequest() {
		const response = await fetch(`/api/deleteuser/${selectedToDeleteId}`, { method: 'POST' });

		cancelDeleteProcess();
		getUsersRequest();
	}

	async function getUsersRequest() {
		const response = await fetch(`/api/getallusers`, { method: 'GET' });

		const body = await response.json();
		users = body.users;
	}

	// const numberOfColumns: number = 6;
	// const namesOfColumns: string[] = [
	// 	'ID',
	// 	'Vardas',
	// 	'Pavardė',
	// 	'Elektroninis paštas',
	// 	'Vartotojo vardas',
	// 	'Rolė'
	// ];
	const numberOfColumns: number = 2;
	const namesOfColumns: string[] = ['ID', 'Vartotojo vardas'];
</script>

<svelte:head>
	<title>Visi vartotojai - Kavos Skrudykla</title>
</svelte:head>

{#if wantsToAdd}
	<AddUser />
{/if}

{#if wantsToDelete}
	<button
		onclick={cancelDeleteProcess}
		aria-label="Paspauskite, kad baigtumėte ištrynimo procesą."
		id="overlay"
		class="fixed left-0 top-0 z-10 h-full w-full cursor-default bg-black opacity-60"
	></button>
	<div
		id="modal"
		class="fixed left-1/2 top-1/2 z-10 h-fit w-[32rem] -translate-x-1/2 -translate-y-1/2 transform rounded-lg bg-white px-4 pb-3 pt-2"
	>
		<p class="mb-2 mt-1 text-2xl font-medium">Vartotojo ištrynimas</p>
		<p class="mb-5 mt-1">Ar tikrai norite ištrinti vartotoją ({selectedToDeleteUsername})?</p>

		<div class="flex grow justify-end gap-2">
			<button
				onclick={cancelDeleteProcess}
				class="10 rounded border border-gray-600 bg-white px-6 py-1 font-medium text-gray-600 hover:bg-gray-200"
				>Atšaukti</button
			>
			<button
				onclick={sendDeleteRequest}
				class="rounded bg-red-500 px-6 py-1 font-medium text-white hover:bg-red-600"
				>Ištrinti</button
			>
		</div>
	</div>
{/if}

<div class="flex min-h-screen flex-col">
	<Header userType={USER_ADMIN} />

	<article class="mx-page mb-bottom mt-top grow">
		<div class="mb-8 border-b-2 border-black">
			<p class="mb-2 ml-10 text-4xl font-medium">Vartotojai</p>
		</div>

		<div class="flex justify-end">
			<button
				onclick={startAddUserProcess}
				class="mb-5 mr-2 w-48 self-end rounded-md bg-green-500 py-1 font-medium text-white hover:bg-green-600"
			>
				Pridėti vartotoją</button
			>
		</div>

		<div class="flex flex-col">
			<header
				class="text-md mb-2 grid h-12 grid-cols-2 items-center rounded-md border border-gray-300 bg-white text-center font-medium"
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
				{#if users.length !== 0}
					{#each users as user, i}
						<button
							onclick={() => startDeleteProcess(user.Id, user.username)}
							class="grid h-fit grid-cols-2 items-center border-b border-l border-r border-black text-center hover:bg-gray-400 hover:text-white {i +
								1 ===
							users.length
								? 'rounded-b-md'
								: ''} {i % 2 === 0 ? 'bg-gray-200' : 'bg-gray-300'} {i === 0
								? 'rounded-t-md border-t'
								: ''}"
						>
							<p class="h-full content-center break-words border-r border-black px-2 py-1">
								{user.Id}
							</p>
							<!-- <p class="h-full content-center break-words border-r border-black px-2 py-1">
								{user.name}
							</p>
							<p class="h-full content-center break-words border-r border-black px-2 py-1">
								{user.surname}
							</p>
							<p class="h-full content-center break-words border-r border-black px-2 py-1">
								{user.email}
							</p> -->
							<p class="h-full content-center break-words border-black px-2 py-1">
								{user.username}
							</p>
							<!-- <p class="h-full content-center break-words border-black px-2 py-1">
								{user.role}
							</p> -->
						</button>
					{/each}
				{:else}
					<div
						class="grid h-10 grid-cols-1 items-center rounded-md border border-gray-300 bg-white text-center"
					>
						<p class="text-sm font-medium">Vartotojų nėra</p>
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
