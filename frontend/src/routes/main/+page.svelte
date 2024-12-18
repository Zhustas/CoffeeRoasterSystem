<script lang="ts">
	import type { PageData } from './$types';
	import type { Coffee } from '$lib/coffee';
	import { USER_ADMIN, USER_CUSTOMER, USER_ROASTER } from '$lib/constants/UserTypeConstants';
	import MainCustomer from '$lib/components/MainCustomer.svelte';
	import MainAdmin from '$lib/components/MainAdmin.svelte';
	import MainRoaster from '$lib/components/MainRoaster.svelte';

	let { data }: { data: PageData } = $props();
	let user = data.user;
	let coffees: Coffee[] = data.coffees;
</script>

<svelte:head>
	<title>Pagrindinis - Kavos Skrudykla</title>
</svelte:head>

{#if user.role === USER_CUSTOMER}
	<MainCustomer {coffees} />
{:else if user.role === USER_ADMIN}
	<MainAdmin {user} />
{:else if user.role === USER_ROASTER}
	<MainRoaster {user} />
{/if}
