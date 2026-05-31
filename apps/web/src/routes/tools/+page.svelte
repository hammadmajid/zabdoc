<script lang="ts">
	import SEO from "$lib/components/seo.svelte";
	import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "$lib/components/ui/card";
	import { fade } from "svelte/transition";
	import Code from "@lucide/svelte/icons/code";
	import ExternalLink from "@lucide/svelte/icons/external-link";

	type Prompt = {
		id: string;
		title: string;
		description: string;
		content: string;
	};

	type GPT = {
		id: string;
		name: string;
		description: string;
		url: string;
		icon: any;
	};

	const prompts: Prompt[] = [
		{
			id: "cpp-general",
			title: "C++ Instructions General",
			description: "General C++ programming guidelines and best practices",
			content: "Provide comprehensive C++ programming instructions covering syntax, features, and conventions for general purpose programming."
		},
		{
			id: "mids-proposal",
			title: "Mids Project Proposal",
			description: "Template for mid-semester project proposals",
			content: "Generate a structured mid-semester project proposal template including objectives, timeline, deliverables, and evaluation criteria."
		},
		{
			id: "final-proposal",
			title: "Final Project Proposal",
			description: "Template for final project proposals",
			content: "Create a comprehensive final project proposal template with detailed project description, requirements, implementation plan, and success metrics."
		},
		{
			id: "cpp-lab",
			title: "C++ Instructions Lab Task",
			description: "C++ instructions for laboratory tasks and experiments",
			content: "Provide detailed C++ lab task instructions including experiment objectives, procedure, expected outputs, and lab report guidelines."
		},
		{
			id: "cpp-linux",
			title: "C++ Instructions Linux",
			description: "C++ development on Linux systems",
			content: "Guide for C++ programming on Linux including compilation, debugging tools, GCC/Clang usage, and Linux-specific libraries."
		},
		{
			id: "python",
			title: "Python Instructions",
			description: "Python programming guidelines and practices",
			content: "Comprehensive Python programming instructions covering syntax, paradigms, libraries, and best practices for different use cases."
		}
	];

	const gpts: GPT[] = [
		{
			id: "zab-cpp",
			name: "Zab_C++",
			description: "Writes C++ code according to SZABIST standards with explanatory comments",
			url: "https://chatgpt.com/g/g-6a03af8b65a481918de130f20c7f6dc0-zab-c",
			icon: Code
		}
	];

	let selectedPrompt = $state<Prompt | null>(null);
	let textareaValue = $state("");

	function selectPrompt(prompt: Prompt) {
		selectedPrompt = prompt;
		textareaValue = prompt.content;
	}

	function copyToClipboard() {
		navigator.clipboard.writeText(textareaValue);
	}
</script>

<SEO
	title="Tools - zabdoc"
	description="AI prompts and tools for SZABIST students."
	canonical="https://zabdoc.xyz/tools"
	url="https://zabdoc.xyz/tools"
/>

<div class="flex min-h-screen flex-col">
	<!-- Header -->
	<div class="mb-8 text-center" in:fade={{ duration: 300 }}>
		<div class="neo-border neo-shadow-lg mb-6 inline-block -rotate-2 bg-primary px-8 py-4">
			<h1 class="text-4xl font-black tracking-tight uppercase md:text-5xl">Tools</h1>
		</div>
		<p class="mx-auto max-w-xl text-lg font-medium">
			Helpful prompts and utilities for your assignments and projects.
		</p>
	</div>

	<!-- GPTs Section -->
	<div class="px-4 py-8 border-b-2 border-border">
		<div class="mb-6">
			<h2 class="text-2xl font-black uppercase mb-4 ml-4">Custom GPTs</h2>
		</div>
		<div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 px-4">
			{#each gpts as gpt (gpt.id)}
				<a
					href={gpt.url}
					target="_blank"
					rel="noopener noreferrer"
					class="neo-border neo-shadow group flex flex-col items-center gap-4 bg-card p-6 text-center transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none"
					in:fade={{ duration: 300 }}
				>
					<div class="neo-border neo-shadow-sm bg-primary p-4 group-hover:bg-secondary transition-colors">
						<svelte:component this={gpt.icon} class="size-10 text-primary-foreground" />
					</div>
					<div>
						<h3 class="font-bold text-lg uppercase">{gpt.name}</h3>
						<p class="text-sm text-muted-foreground mt-2">{gpt.description}</p>
					</div>
					<div class="mt-auto pt-4 flex items-center gap-2 text-sm font-semibold text-primary">
						Open GPT
						<ExternalLink class="size-4" />
					</div>
				</a>
			{/each}
		</div>
	</div>

	<!-- Content -->
	<div class="flex flex-1 gap-8 px-4 py-8">
		<!-- Prompts List -->
		<div class="w-full md:w-1/3">
			<div class="space-y-3">
				{#each prompts as prompt (prompt.id)}
					<button
						onclick={() => selectPrompt(prompt)}
						class={`neo-border neo-shadow w-full px-4 py-3 text-left transition-all ${
							selectedPrompt?.id === prompt.id
								? "bg-primary text-primary-foreground"
								: "bg-card hover:translate-x-1 hover:translate-y-1 hover:shadow-none"
						}`}
					>
						<div class="font-bold">{prompt.title}</div>
						<div class="text-sm opacity-75">{prompt.description}</div>
					</button>
				{/each}
			</div>
		</div>

		<!-- Textarea -->
		<div class="w-full md:w-2/3">
			{#if selectedPrompt}
				<div in:fade={{ duration: 200 }}>
					<Card class="flex flex-col h-full">
						<CardHeader>
							<CardTitle>{selectedPrompt.title}</CardTitle>
							<CardDescription>{selectedPrompt.description}</CardDescription>
						</CardHeader>
						<CardContent class="flex flex-1 flex-col gap-4">
							<textarea
								bind:value={textareaValue}
								class="neo-border-sm neo-shadow-sm flex-1 w-full bg-background px-4 py-3 text-sm font-medium transition-all outline-none selection:bg-primary selection:text-primary-foreground placeholder:text-muted-foreground focus:translate-x-[2px] focus:translate-y-[2px] focus:shadow-none disabled:cursor-not-allowed disabled:opacity-50 resize-none"
								placeholder="Your prompt will appear here..."
								rows="15"
							/>
							<button
								onclick={copyToClipboard}
								class="neo-border neo-shadow w-full bg-secondary px-4 py-2 font-bold text-secondary-foreground uppercase transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none"
							>
								Copy to Clipboard
							</button>
						</CardContent>
					</Card>
				</div>
			{:else}
				<div class="flex items-center justify-center rounded-lg border-2 border-dashed border-muted bg-muted/50 p-8 text-center">
					<p class="text-muted-foreground">Select a prompt to get started</p>
				</div>
			{/if}
		</div>
	</div>
</div>
