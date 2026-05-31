<script lang="ts">
	import SEO from "$lib/components/seo.svelte";
	import Database from "@lucide/svelte/icons/database";
	import FileText from "@lucide/svelte/icons/file-text";
	import Lightbulb from "@lucide/svelte/icons/lightbulb";
	import { fade } from "svelte/transition";
</script>

<SEO
	title="zabdoc"
	description="Generate assignment and lab task PDFs for SZABIST students."
	canonical="https://zabdoc.xyz/"
	url="https://zabdoc.xyz/"
/>

<div class="flex min-h-screen flex-col lg:flex-row">
	<!-- Left Section: Content & Actions -->
	<div class="flex flex-1 flex-col justify-center px-8 py-16 lg:px-12">
		<div in:fade={{ duration: 300 }}>
			<!-- Title -->
			<div class="neo-border neo-shadow-lg mb-8 -rotate-2 bg-primary px-8 py-4 inline-block">
				<h1 class="text-4xl font-black tracking-tight uppercase md:text-6xl">zabdoc</h1>
			</div>

			<!-- Subtitle -->
			<p class="mb-4 text-lg font-medium md:text-xl">
				utility tools for
				<span class="neo-border-sm bg-secondary px-2 py-0.5 font-bold">*SZABIST</span>
				students.
			</p>

			<!-- Description -->
			<p class="mb-12 max-w-md text-base font-medium text-muted-foreground">
				Choose an option to get started with document generation, data scraping, or AI-powered tools.
			</p>

			<!-- Buttons Grid -->
			<div class="grid w-full max-w-2xl grid-cols-2 gap-6 auto-rows-max mb-12">
				<a
					href="/scrape"
					class="neo-border neo-shadow flex cursor-pointer items-center gap-3 bg-accent px-8 py-6 text-lg font-black tracking-wide text-secondary-foreground uppercase transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none"
				>
					<Database class="size-7" />
					<span>Scrape<br />Data</span>
				</a>
				<a
					href="/document"
					class="neo-border neo-shadow flex cursor-pointer items-center gap-3 bg-secondary px-8 py-6 text-lg font-black tracking-wide text-accent-foreground uppercase transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none"
				>
					<FileText class="size-7" />
					<span>Generate<br />Document</span>
				</a>
				<a
					href="/tools"
					class="neo-border neo-shadow flex cursor-pointer items-center gap-3 bg-primary px-8 py-6 text-lg font-black tracking-wide text-primary-foreground uppercase transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none col-span-2"
				>
					<Lightbulb class="size-7" />
					<span>Tools</span>
				</a>
			</div>

			<!-- Disclaimer -->
			<p class="neo-border-sm inline-block bg-muted px-4 py-2 text-sm font-medium max-w-md">
				*this project is <em>not sponsored, affiliated, endorsed, or approved by</em> SZABIST.
				<a href="/about#disclaimer" class="font-bold underline hover:text-primary">Learn more.</a>
			</p>
		</div>
	</div>

	<!-- Right Section: Graphic/Visual -->
	<div class="flex flex-1 items-center justify-center px-8 py-16 lg:px-12 bg-muted/30" in:fade={{ duration: 300, delay: 100 }}>
		<div class="relative w-full max-w-sm">
			<!-- Decorative SVG Graphic -->
			<svg
				class="w-full h-auto"
				viewBox="0 0 400 400"
				xmlns="http://www.w3.org/2000/svg"
			>
				<!-- Background shapes -->
				<rect
					x="20"
					y="20"
					width="160"
					height="160"
					fill="none"
					stroke="currentColor"
					stroke-width="3"
					class="text-primary opacity-20"
				/>
				<rect
					x="220"
					y="20"
					width="160"
					height="160"
					fill="none"
					stroke="currentColor"
					stroke-width="3"
					class="text-secondary opacity-20"
					transform="rotate(45 300 100)"
				/>
				<rect
					x="120"
					y="220"
					width="160"
					height="160"
					fill="none"
					stroke="currentColor"
					stroke-width="3"
					class="text-accent opacity-20"
				/>

				<!-- Central circles -->
				<circle cx="200" cy="200" r="80" fill="none" stroke="currentColor" stroke-width="2" class="text-primary opacity-30" />
				<circle cx="200" cy="200" r="60" fill="none" stroke="currentColor" stroke-width="2" class="text-secondary opacity-30" />
				<circle cx="200" cy="200" r="40" fill="none" stroke="currentColor" stroke-width="2" class="text-accent opacity-30" />

				<!-- Accent lines -->
				<line x1="200" y1="140" x2="200" y2="60" stroke="currentColor" stroke-width="3" class="text-primary" />
				<line x1="260" y1="200" x2="340" y2="200" stroke="currentColor" stroke-width="3" class="text-secondary" />
				<line x1="200" y1="260" x2="200" y2="340" stroke="currentColor" stroke-width="3" class="text-accent" />
				<line x1="140" y1="200" x2="60" y2="200" stroke="currentColor" stroke-width="3" class="text-primary opacity-60" />

				<!-- Corner accents -->
				<rect x="30" y="30" width="30" height="30" fill="none" stroke="currentColor" stroke-width="2" class="text-primary" />
				<rect x="340" y="30" width="30" height="30" fill="none" stroke="currentColor" stroke-width="2" class="text-secondary" />
				<rect x="30" y="340" width="30" height="30" fill="none" stroke="currentColor" stroke-width="2" class="text-accent" />
				<rect x="340" y="340" width="30" height="30" fill="none" stroke="currentColor" stroke-width="2" class="text-primary opacity-60" />
			</svg>

			<!-- Floating cards effect -->
			<div
				class="neo-border neo-shadow absolute -top-4 -right-4 bg-primary px-4 py-2 text-xs font-bold text-primary-foreground uppercase w-24 h-24 flex items-center justify-center text-center"
				style="animation: float 6s ease-in-out infinite;"
			>
				Scrape
			</div>
			<div
				class="neo-border neo-shadow absolute bottom-8 -left-6 bg-secondary px-4 py-2 text-xs font-bold text-accent-foreground uppercase w-24 h-24 flex items-center justify-center text-center"
				style="animation: float 5s ease-in-out infinite 1s;"
			>
				Generate
			</div>
			<div
				class="neo-border neo-shadow absolute -bottom-4 -right-8 bg-accent px-4 py-2 text-xs font-bold text-secondary-foreground uppercase w-24 h-24 flex items-center justify-center text-center"
				style="animation: float 7s ease-in-out infinite 0.5s;"
			>
				Tools
			</div>
		</div>
	</div>
</div>

<style>
	@keyframes float {
		0%,
		100% {
			transform: translateY(0px);
		}
		50% {
			transform: translateY(-20px);
		}
	}
</style>
