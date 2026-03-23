<script lang="ts">
	import { onNavigate } from "$app/navigation";
	import { page } from "$app/stores";
	import favicon from "$lib/assets/favicon.svg";
	import Settings from "@lucide/svelte/icons/settings";
	import Github from "@lucide/svelte/icons/github";
	import { buttonVariants } from "$lib/components/ui/button";
	import { cn } from "$lib/utils";
	import { wizardStore } from "$lib/stores/wizard-store.svelte";
	import { changelog } from "$lib/changelog";
	import "../app.css";

	onNavigate((navigation) => {
		if (!document.startViewTransition) return;

		return new Promise((resolve) => {
			document.startViewTransition(async () => {
				resolve();
				await navigation.complete;
			});
		});
	});

	function handleZabdocClick() {
		wizardStore.reset();
	}

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="mx-auto flex min-h-screen max-w-5xl flex-col px-4">
	<header class="py-6">
		<div class="flex items-center justify-between">
			<nav class="flex items-center gap-4">
				<a
					href="/"
					onclick={handleZabdocClick}
					class="neo-border-sm neo-shadow-sm bg-primary px-4 py-2 text-xl font-black tracking-tight uppercase transition-all hover:translate-x-0.5 hover:translate-y-0.5 hover:shadow-none"
				>
					zabdoc
				</a>
				<a
					href="/changelog"
					class="neo-border-sm bg-muted px-2 py-0.5 text-xs font-bold transition-colors hover:bg-primary"
				>
					{changelog[0].version}
				</a>
			</nav>
			<nav class="flex items-center gap-2">
				<a
					href="/about"
					class={cn(
						"px-3 py-1 text-sm font-bold tracking-wide uppercase transition-colors hover:text-primary",
						$page.url.pathname === "/about" && "underline"
					)}
				>
					about
				</a>
				<a href="/settings" class={buttonVariants({ variant: "ghost", size: "icon" })}>
					<Settings class="size-5" />
				</a>
			</nav>
		</div>
	</header>

	<main class="flex-1 py-6">
		{@render children?.()}
	</main>

	<footer class="mt-4 py-6">
		<div class="flex items-center justify-between">
			<div class="flex items-center gap-3">
				<a
					href="https://github.com/hammadmajid/zabdoc"
					target="_blank"
					rel="noreferrer"
					class="neo-border-sm bg-card p-2 transition-colors hover:bg-muted"
				>
					<Github class="size-5" />
				</a>
				<span class="text-sm font-medium text-muted-foreground"> Open Source </span>
			</div>
			<div class="text-sm font-medium">
				<a
					href="https://github.com/hammadmajid/zabdoc/blob/main/LICENSE"
					target="_blank"
					rel="noreferrer"
					class="text-muted-foreground transition-colors hover:text-foreground"
				>
					GPL v3
				</a>
				<span class="mx-2 text-muted-foreground">•</span>
				<span class="neo-border-sm bg-muted px-2 py-0.5 text-xs font-bold"> 2026 </span>
			</div>
		</div>
	</footer>
</div>
