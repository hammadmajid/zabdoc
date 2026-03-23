<script lang="ts">
	import SEO from "$lib/components/seo.svelte";
	import Database from "@lucide/svelte/icons/database";
	import FileText from "@lucide/svelte/icons/file-text";
	import { fade } from "svelte/transition";
</script>

<SEO
	title="zabdoc"
	description="Generate assignment and lab task PDFs for SZABIST students."
	canonical="https://zabdoc.xyz/"
	url="https://zabdoc.xyz/"
/>

<div class="flex min-h-[70vh] flex-col">
	<!-- Hero Section -->
	<div class="mb-4 text-center" in:fade={{ duration: 300 }}>
		<div class="neo-border neo-shadow-lg mb-6 inline-block -rotate-2 bg-primary px-8 py-4">
			<h1 class="text-4xl font-black tracking-tight uppercase md:text-6xl">zabdoc</h1>
		</div>
		<p class="mx-auto max-w-xl text-lg font-medium">
			utility tools for
			<span class="neo-border-sm bg-secondary px-2 py-0.5 font-bold">*SZABIST</span>
			students.
		</p>
	</div>

	<!-- Choice Section -->
	<div class="flex flex-1 flex-col items-center justify-center px-4 py-8 text-center">
		<div class="neo-border neo-shadow-lg mb-8 -rotate-2 bg-primary px-8 py-4">
			<h2 class="text-4xl font-black tracking-tight uppercase md:text-5xl">
				What do you need?
			</h2>
		</div>

		<p class="mb-8 max-w-md text-lg font-medium">Choose an option to get started.</p>

		<div class="flex flex-col gap-6 sm:flex-row">
			<a
				href="/scrape"
				class="neo-border neo-shadow flex cursor-pointer items-center gap-3 bg-accent px-10 py-6 text-xl font-black tracking-wide text-secondary-foreground uppercase transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none"
			>
				<Database class="size-8" />
				Scrape Data
			</a>
			<a
				href="/document"
				class="neo-border neo-shadow flex cursor-pointer items-center gap-3 bg-secondary px-10 py-6 text-xl font-black tracking-wide text-accent-foreground uppercase transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none"
			>
				<FileText class="size-8" />
				Generate Document
			</a>
		</div>
	</div>

	<!-- Disclaimer -->
	<div class="mt-4 text-center" in:fade={{ duration: 300, delay: 200 }}>
		<p class="neo-border-sm inline-block max-w-md bg-muted px-4 py-2 text-sm font-medium">
			*this project is <em>not sponsored, affiliated, endorsed, or approved by</em> SZABIST.
			<a href="/about#disclaimer" class="font-bold underline hover:text-primary"
				>Learn more.</a
			>
		</p>
	</div>
</div>
