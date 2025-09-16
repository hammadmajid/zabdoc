<script lang="ts">
    import { page } from "$app/stores";
    import { buttonVariants } from "$lib/components/ui/button";
    import AlertTriangle from "@lucide/svelte/icons/alert-triangle";
    import Home from "@lucide/svelte/icons/home";
    import RotateCcw from "@lucide/svelte/icons/rotate-ccw";
</script>

<svelte:head>
    <title>error | zabdoc</title>
</svelte:head>

<div class="flex flex-col items-center justify-center text-center h-full px-4">
    <div class="mb-8">
        <AlertTriangle class="mx-auto h-16 w-16 text-destructive mb-4" />
        <h1 class="text-5xl font-bold tracking-tight">Oops!</h1>
        <p class="mt-4 text-lg text-muted-foreground max-w-xl">
            Something went wrong while processing your request.
            <span class="border-b border-dotted border-foreground/50"
                >zabdoc</span
            >
            encountered an unexpected error.
        </p>
    </div>

    <div class="mb-8 p-4 bg-muted rounded-lg max-w-md">
        <p class="text-sm text-muted-foreground font-mono">
            Error {$page.status}: {$page.error?.message ||
                "Unknown error occurred"}
        </p>
    </div>

    <div class="flex gap-4">
        <button
            onclick={() => window.location.reload()}
            class={buttonVariants({ variant: "outline", size: "lg" })}
        >
            <RotateCcw class="mr-2 h-4 w-4" />
            Try Again
        </button>
        <a href="/" class={buttonVariants({ variant: "default", size: "lg" })}>
            <Home class="mr-2 h-4 w-4" />
            Go Home
        </a>
    </div>

    <div class="mt-8">
        <p class="text-sm text-muted-foreground">
            If this problem persists, please{" "}
            <a
                href="https://github.com/hammadmajid/zabdoc/issues"
                target="_blank"
                rel="noopener noreferrer"
                class="underline hover:no-underline"
            >
                report an issue
            </a>
            {" "}on GitHub.
        </p>
    </div>
</div>
