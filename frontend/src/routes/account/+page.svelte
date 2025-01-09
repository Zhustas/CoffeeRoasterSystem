<script lang="ts">
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import Footer from '$lib/components/Footer.svelte';
	import Header from '$lib/components/Header.svelte';
	import { USER_ADMIN, USER_CUSTOMER, USER_ROASTER } from '$lib/constants/UserTypeConstants';

	let { data }: { data: PageData } = $props();
	let user = data;

	onMount(() => {
		name = user.name;
		surname = user.surname;
		email = user.email;
		username = user.username;
		password = 'xxxxxxxx';
	});

	// Laukai
	let name: string = $state();
	let surname: string = $state();
	let email: string = $state();
	let username: string = $state();
	let password: string = $state();

	function getRoleLT(role: string) {
		if (role === 'customer') return 'Klientas';
		if (role === 'roaster') return 'Darbuotojas';
		if (role === 'admin') return 'Administratorius';
	}

	async function logOut() {
		const response = await fetch('/api/logout', {
			method: 'GET'
		});

		document.location.assign('/');
	}
</script>

<svelte:head>
	<title>Paskyra - Kavos Skrudykla</title>
</svelte:head>

<div class="flex min-h-screen flex-col">
	{#if user.role === USER_CUSTOMER}
		<Header userType={USER_CUSTOMER} />
	{:else if user.role === USER_ROASTER}
		<Header userType={USER_ROASTER} />
	{:else if user.role === USER_ADMIN}
		<Header userType={USER_ADMIN} />
	{/if}

	<article class="mx-page mt-top mb-bottom flex grow flex-col">
		<div class="mb-8 border-b-2 border-black">
			<p class="mb-2 ml-10 text-4xl font-medium">Mano paskyra</p>
		</div>

		<div class="flex grow gap-5">
			<div class="w-72 rounded-md bg-white">
				<p class="mb-1 mt-2 text-center text-xl font-medium">{name} {surname}</p>
				<p class="mb-1 text-center text-sm text-gray-500">ID: {user.Id}</p>
				<p class="text-center text-sm text-gray-500">{getRoleLT(user.role)}</p>
			</div>

			<div class="flex grow flex-col gap-3 rounded-md bg-white p-5">
				<div class="flex flex-col">
					<label for="name" class="mb-1 ml-1">Vardas</label>
					<input
						id="name"
						bind:value={name}
						type="text"
						disabled
						class="rounded-md border-gray-500 text-gray-500"
					/>
				</div>

				<div class="flex flex-col">
					<label for="surname" class="mb-1 ml-1">Pavardė</label>
					<input
						id="surname"
						bind:value={surname}
						type="text"
						disabled
						class="rounded-md border-gray-500 text-gray-500"
					/>
				</div>

				<div class="flex flex-col">
					<label for="email" class="mb-1 ml-1">Elektroninis paštas</label>
					<input
						id="email"
						bind:value={email}
						type="email"
						disabled
						class="rounded-md border-gray-500 text-gray-500"
					/>
				</div>

				<div class="flex flex-col">
					<label for="username" class="mb-1 ml-1">Vartotojo vardas</label>
					<input
						id="username"
						bind:value={username}
						type="text"
						disabled
						class="rounded-md border-gray-500 text-gray-500"
					/>
				</div>

				<div class="mb-5 flex flex-col">
					<label for="password" class="mb-1 ml-1">Slaptažodis</label>
					<input
						id="password"
						bind:value={password}
						type="password"
						disabled
						class="rounded-md border-gray-500 text-gray-500"
					/>
				</div>

				<div class="flex justify-between">
					<button
						onclick={logOut}
						class="w-32 rounded border border-gray-600 bg-white px-6 py-1 font-medium text-gray-600 hover:bg-gray-200"
						>Atsijungti</button
					>
					<button
						class="w-52 self-end rounded border border-yellow-500 bg-yellow-500 px-6 py-1 font-medium text-white hover:border-yellow-600 hover:bg-yellow-600"
						>Atnaujinti informaciją</button
					>
				</div>
			</div>
		</div>
	</article>

	<Footer />
</div>
