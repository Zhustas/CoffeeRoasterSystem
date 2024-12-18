<script lang="ts">
	import {
		USER_ADMIN,
		USER_CUSTOMER,
		USER_NONE,
		USER_ROASTER
	} from '$lib/constants/UserTypeConstants';

	let { userType }: { userType: string } = $props();
</script>

<header class="fixed w-full">
	<nav class="flex h-16 border-b border-gray-400 bg-white px-64">
		<a href="/main" class="self-center p-2">
			<img alt="Kavos Skrudykla" src="coffee_roaster_title.png" class="h-12" />
		</a>

		<div class="flex grow items-center justify-end font-medium">
			{#if userType !== USER_NONE}
				{#if userType === USER_CUSTOMER}
					<a href="/subscription" class="mx-4 p-2">Prenumerata</a>
					<a href="/history" class="mx-4 p-2">Užsakymai</a>
				{:else}
					{#if userType === USER_ADMIN}
						<a href="/users" class="mx-4 p-2">Vartotojai</a>
						<a href="/orders" class="mx-4 p-2">Visi užsakymai</a>
					{/if}
					{#if userType === USER_ROASTER}
						<a href="/inventory" class="mx-4 p-2">Kavos inventorius</a>
						<a href="/orders" class="mx-4 p-2">Visi užsakymai</a>
						<a href="/neworders" class="mx-4 p-2">Nauji užsakymai</a>
					{/if}
				{/if}
				<a href="/account" class="mx-4 p-2">Paskyra</a>
				<div class="h-full w-px bg-gray-400"></div>
				<a href="/faq" class="mx-4 p-2">DUK</a>
			{/if}
		</div>
	</nav>
</header>
