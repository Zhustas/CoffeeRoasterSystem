<script lang="ts">
	import type { Coffee } from '../../../lib/Coffee';

	const coffee: Coffee = {
		ID: 1,
		description:
			'Kava - gėrimas, gaminamas iš skrudintų ir maltų Afrikos kilmės atogrąžų visžalių kavamedžių sėklų. Kava yra vienas iš trijų populiariausių gėrimų pasaulyje (kartu su vandeniu ir arbata) ir viena pelningiausių tarptautinių prekių. Nors iš kavos gaminama begalė gėrimų, įskaitant espreso, kapučino, moka ir latte, jos populiarumą daugiausia lemia žvalumo suteikianti kava, kurią sukelia kavoje esantis alkaloidas kofeinas.',
		name: 'Kava',
		price: 15.99,
		roast_type: 'light',
		stock: 10
	};
</script>

<div class="flex min-h-screen flex-col">
	<header>
		<nav class="flex h-16 border-b border-gray-400 bg-white px-64">
			<a aria-label="Perkrauti šį puslapį" href="/main" class="self-center p-2">
				<img alt="Tekstas - Kavos Skrudykla" src="../coffee_roaster_title.png" class="h-12" />
			</a>
			<div class="flex grow items-center justify-end font-medium">
				<a href="/main" class="mx-4 p-2">Prenumerata</a>
				<a href="/main" class="mx-4 p-2">Užsakymai</a>
				<a href="/main" class="mx-4 p-2">Paskyra</a>
				<div class="h-full w-px bg-gray-400"></div>
				<a href="/main" class="mx-4 p-2">DUK</a>
			</div>
		</nav>
	</header>
	<article class="mx-64 flex grow">
		<div id="image" class="flex basis-1/2 justify-center px-5 pb-20 pt-10">
			<div class="rounded-md bg-white p-5">
				<img alt="Coffee bean" src="../coffee_bean.jpg" class="h-full" />
			</div>
		</div>
		<div id="information" class="flex basis-1/2 flex-col bg-white px-5 pb-20 pt-10">
			<p class="mb-2 text-4xl font-medium">{coffee.name}</p>
			<div class="mb-8 h-44 rounded-md bg-gray-100 px-2 py-1">
				<p class="text-sm">{coffee.description}</p>
			</div>
			<p class="mb-2">Tipas: <span class="font-medium">tamsiai skrudinta</span></p>

			<div class="mb-8 flex">
				<label for="roast_type" class="mr-2 self-center">Pasirinkite kiekį</label>
				<select id="roast_type" class="mr-1 rounded-md">
					<option>100</option>
					<option>200</option>
					<option>400</option>
					<option>500</option>
					<option>750</option>
					<option>1000</option>
				</select>
				<p class="self-center font-medium">g.</p>
			</div>
			<div class="mb-14 flex justify-end">
				<p class=" mr-2 text-3xl font-medium">15,00 €</p>
			</div>

			<div class="flex grow flex-col justify-end">
				<button
					class="bg-button hover:bg-hover-button h-20 w-full rounded-md text-2xl font-medium tracking-wider"
					>PIRKTI</button
				>
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
