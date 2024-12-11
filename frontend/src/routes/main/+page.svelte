<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import type { Coffee } from '$lib/Coffee';
	import type { PageData } from './$types';
	import Footer from '../../components/Footer.svelte';
	import CoffeeTab from '../../components/CoffeeTab.svelte';

	let { data }: { data: PageData } = $props();
	let coffees: Coffee[] = data.coffees;

	let intersecting = $state(false);
	let header: HTMLElement;
	onMount(() => {
		if (typeof IntersectionObserver !== 'undefined') {
			const headerObserver = new IntersectionObserver(
				(entries) => {
					intersecting = entries[0].isIntersecting;
				},
				{
					root: null,
					threshold: 0,
					rootMargin: `${header.getBoundingClientRect().height}px`
				}
			);

			headerObserver.observe(header);
		}

		slides = slider.querySelectorAll('.slide');
		goToSlide(0);
		removeHiddenSlides();

		myInterval = setInterval(myIntervalFunction, 5000);
	});

	onDestroy(() => {
		clearInterval(myInterval);
	});

	let hovering = $state(false);

	function changeHoveringState(state: boolean) {
		hovering = state;
	}

	let myInterval: number;
	function myIntervalFunction() {
		if (!hovering) {
			slideRight();
		}
	}

	$effect(() => {
		dots = dotsContainer.children;
	});

	// Slide efektas
	let slider: HTMLElement;
	let currentSlide = 0;
	let slides: NodeListOf<Element> = $state();
	let dots: HTMLCollection;
	let dotsContainer: HTMLElement;

	function removeHiddenSlides() {
		slides.forEach((s, i) => {
			s.classList.remove('hidden');
		});
	}

	function goToSlide(slide: number) {
		slides.forEach((s, i) => {
			s.style.transform = `translateX(${100 * (i - slide)}%)`;
		});
	}

	function slideLeft() {
		currentSlide--;

		if (currentSlide === -1) {
			currentSlide = slides.length - 1;
		}

		goToSlide(currentSlide);
		activateDot(currentSlide);
	}

	function slideRight() {
		currentSlide++;

		if (currentSlide === slides.length) {
			currentSlide = 0;
		}

		goToSlide(currentSlide);
		activateDot(currentSlide);
	}

	function dotContainerOnClick(event: MouseEvent) {
		const target = event.target;
		if (target.dataset.slide) {
			handleDotClick(Number(target.dataset.slide) - 1);
		}
	}

	function handleDotClick(slide: number) {
		currentSlide = slide;
		goToSlide(slide);
		activateDot(slide);
	}

	function activateDot(slide: number) {
		for (let dot of dots) {
			dot.classList.remove('bg-gray-400', 'bg-white');
			if (Number(dot.dataset.slide) === slide + 1) {
				dot.classList.add('bg-gray-400');
			} else {
				dot.classList.add('bg-white');
			}
		}
	}

	// Filtrai
	const FILTER_LOWEST = 'lowest';
	const FILTER_HIGHEST = 'highest';

	let selectedFilter = $state();

	function filterCoffees() {
		if (selectedFilter === FILTER_LOWEST) {
			coffees.sort((a, b) => a.price - b.price);
		} else if (selectedFilter === FILTER_HIGHEST) {
			coffees.sort((a, b) => b.price - a.price);
		}

		if (searchText) {
			const regex = new RegExp(`${searchText.toLocaleLowerCase()}`);

			const filtered = coffees.filter((coffee) => coffee.name.toLocaleLowerCase().match(regex));

			return filtered;
		}

		return coffees;
	}

	let searchText: string = $state();

	// Rodomos kavos
	interface Range {
		from: number;
		to: number;
	}

	const coffeesRange: Range = { from: 0, to: 30 };
	let currentRangeIndex = 0;
</script>

<div class="flex min-h-screen flex-col">
	<header bind:this={header}>
		<nav
			class="flex h-16 border-b border-gray-400 bg-white px-64 {!intersecting
				? 'fixed top-0 z-50 w-full opacity-95'
				: ''}"
		>
			<a aria-label="Perkrauti šį puslapį" href="/main" class="self-center p-2">
				<img alt="Tekstas - Kavos Skrudykla" src="coffee_roaster_title.png" class="h-12" />
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

	<article class="mx-64 mb-10 mt-5 flex grow flex-col">
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			onmouseenter={() => changeHoveringState(true)}
			onmouseleave={() => changeHoveringState(false)}
			id="recommendations"
			class="mx-10 flex flex-col items-center"
		>
			<div id="top" class="flex justify-center gap-10">
				<div class="z-10 flex size-12 self-center rounded-full bg-white text-center text-3xl">
					<button onclick={slideLeft} class="grow rounded-full hover:bg-gray-200">&larr;</button>
				</div>
				<div bind:this={slider} class="relative h-72 w-[52rem] overflow-hidden rounded-md bg-white">
					<div class="slide absolute h-full w-full transition-transform duration-1000">
						<div class="h-full px-4 py-2">
							<h3 class="mb-3 text-2xl font-medium">Ką daro skrudinimo procesas?</h3>
							<p class="mb-1">
								Kavos pupelė - tai kavos vyšnios viduje esanti sėkla. Prieš skrudinant kavos pupelės
								yra žalios ir beveik neturi aromato, išskyrus žemės, žolės kvapą. Skrudinimo proceso
								metu iš kavos pupelių sukuriamas skanus kavos puodelis, kurį geriate.
							</p>
							<p class="mb-1">
								Skrudinimo metu kavos pupelės skrudinamos, todėl jų spalva patamsėja, o skonis tampa
								šokoladinis ir karamelizuotas.
							</p>
							<p>Mūsų parduotuvė parduota trejų tipų skrudintas pupeles:</p>
							<ul class="list-inside list-disc">
								<li class="ml-4">Šviesiai skrudintas pupeles;</li>
								<li class="ml-4">Vidutiniškai skrudintas pupeles;</li>
								<li class="ml-4">Tamsiai skrudintas pupeles.</li>
							</ul>
						</div>
					</div>
					<div class="slide absolute hidden h-full w-full transition-transform duration-1000">
						<div class="flex h-full px-4 py-2">
							<div class="basis-1/2 pr-4">
								<h3 class="mb-3 text-2xl font-medium">Šviesiai skrudintos pupelės</h3>
								<p>
									Kuo ilgiau pupelės skrudinamos, tuo labiau karštis ištraukia kofeiną ir
									rūgštingumą. Tai reiškia, kad šviesiai skrudintose pupelėse yra daugiausia kofeino
									ir daugiausia rūgšties. Lengvo skrudinimo kavos pupelių kilmės skoniai yra labiau
									atpažįstami, nes dėl skrudinimo proceso atsirandantys aromatai dažnai nėra ryškūs.
									Lengvo skrudinimo pupelių rūgštumas dažnai būna lydimas citrusinių vaisių ar
									citrinų atspalvio, kuris kai kuriems žmonėms yra malonus.
								</p>
							</div>
							<div class="flex basis-1/2 justify-center">
								<img
									alt="Šviesiai skrudintos pupelės"
									src="light_roast_coffee_beans.webp"
									class="w-80 self-center"
								/>
							</div>
						</div>
					</div>
					<div class="slide absolute hidden h-full w-full transition-transform duration-1000">
						<div class="flex h-full px-4 py-2">
							<div class="basis-1/2 pr-4">
								<h3 class="mb-3 text-2xl font-medium">Vidutiniškai skrudintos pupelės</h3>
								<p>
									Vidutinio skrudinimo kavos rūšis yra tokia, prie kokios yra pripratęs vidutinis
									amerikietis kavos gėrėjas. Manoma, kad šių skrudintų kavų skonis yra
									subalansuotas. Vidutinio skrudinimo kavos rūgštingumas ir sodrumas gali skirtis,
									bet paprastai būna kažkur per vidurį. Keletas vidutinio skrudinimo kavos
									pavyzdžių: „House blend“, „Breakfast roast“ ir „American Roast“.
								</p>
							</div>
							<div class="flex basis-1/2 justify-center">
								<img
									alt="Vidutiniškai skrudintos pupelės"
									src="medium_roast_coffee_beans.webp"
									class="w-80 self-center"
								/>
							</div>
						</div>
					</div>
					<div class="slide absolute hidden h-full w-full transition-transform duration-1000">
						<div class="flex h-full px-4 py-2">
							<div class="basis-1/2 pr-4">
								<h3 class="mb-3 text-2xl font-medium">Tamsiai skrudintos pupelės</h3>
								<p>
									Tamsiai skrudintos kavos skonis saldesnis, nes kavos pupelėse esantis cukrus spėjo
									karamelizuoti. Ilgesnis skrudinimo procesas padeda kavai įgauti sodresnį skonį ir
									sodrų kūną, todėl ji dažnai būna sviestinė. Be to, iš visų skrudintų kavų jos turi
									mažiausiai rūgštingumo. Tamsiai skrudintoje kavoje yra mažiausiai kofeino, nes ji
									skrudinama ilgiausiai. Prancūziškas skrudinimas laikomas tamsiausiu skrudinimu ir
									pasižymi ryškiu dūmo skoniu.
								</p>
							</div>
							<div class="flex basis-1/2 justify-center">
								<img
									alt="Tamsiai skrudintos pupelės"
									src="dark_roast_coffee_beans.webp"
									class="w-80 self-center"
								/>
							</div>
						</div>
					</div>
				</div>
				<div
					class="z-10 flex size-12 self-center rounded-full bg-white text-center text-3xl hover:bg-gray-200"
				>
					<button onclick={slideRight} class="grow rounded-full">&rarr;</button>
				</div>
			</div>
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<!-- svelte-ignore a11y_click_events_have_key_events -->
			<div
				id="dots"
				bind:this={dotsContainer}
				onclick={dotContainerOnClick}
				class="mt-3 flex grow gap-5"
			>
				{#each slides as _, i}
					<button
						aria-label="Slide dot {i + 1}"
						class="{currentSlide === i
							? 'bg-gray-400'
							: 'bg-white'} size-3 rounded-full border border-gray-300 shadow-lg shadow-black hover:bg-gray-400"
						data-slide={i + 1}
					></button>
				{/each}
			</div>
		</div>

		<hr class="mx-32 my-8 h-0.5 rounded border-0 bg-gray-300" />

		<div id="search" class="flex grow flex-col gap-5">
			<div class="flex h-16 justify-between rounded-md bg-white">
				<div class="ml-2 content-center">
					<select
						bind:value={selectedFilter}
						onselect={filterCoffees}
						class="rounded-md border-2 border-gray-400 hover:border-black focus:border-black focus:outline-none focus:ring-0"
					>
						<option value={FILTER_LOWEST}>Pigiausios viršuje</option>
						<option value={FILTER_HIGHEST}>Brangiausios viršuje</option>
					</select>
				</div>
				<div class="ml-2 mr-2 content-center">
					<input
						bind:value={searchText}
						oninput={filterCoffees}
						placeholder="Ieškoti"
						class="w-64 border-0 border-b-2 border-gray-400 hover:border-black focus:border-black focus:outline-none focus:ring-0"
					/>
				</div>
			</div>

			{#if filterCoffees().length !== 0}
				<div id="products" class="flex flex-wrap gap-5 px-1">
					{#each filterCoffees().slice(coffeesRange.from, coffeesRange.to) as coffee}
						<CoffeeTab id={coffee.ID} name={coffee.name} price={coffee.price} />
					{/each}
				</div>

				<div class="mt-5 flex gap-3 self-center">
					<a href="main/?page=WHERE">
						<div class="flex h-10 w-16 justify-center rounded-md bg-white hover:bg-gray-200">
							<p class="self-center text-3xl">&larr;</p>
						</div>
					</a>
					<div class="flex gap-2">
						{#each Array(Math.ceil(filterCoffees().length / 30)) as num, i}
							{#if currentRangeIndex === i}
								<a href={'FIX'}>
									<div class="flex size-10 justify-center rounded-md bg-gray-300">
										<p class="self-center text-sm">{i + 1}</p>
									</div>
								</a>
							{:else}
								<a href="main/?page={i + 1}">
									<div class="flex size-10 justify-center rounded-md bg-white hover:bg-gray-200">
										<p class="self-center text-sm">{i + 1}</p>
									</div>
								</a>
							{/if}
						{/each}
					</div>
					<a href="main/?page=WHERE">
						<div class="flex h-10 w-16 justify-center rounded-md bg-white hover:bg-gray-200">
							<p class="self-center text-3xl">&rarr;</p>
						</div>
					</a>
				</div>
			{:else}
				<div class="flex grow justify-center">
					<p class="self-center text-sm font-medium">Rezultatų nerasta.</p>
				</div>
			{/if}
		</div>
	</article>
	<Footer />
</div>
