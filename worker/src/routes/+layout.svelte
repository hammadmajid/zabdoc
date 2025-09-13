<script lang="ts">
    import { onNavigate } from "$app/navigation";
    import favicon from "$lib/assets/favicon.svg";
    import Theme from "$lib/components/theme.svelte";
    import { buttonVariants } from "$lib/components/ui/button";
    import { cn } from "$lib/utils";
    import { ModeWatcher } from "mode-watcher";
    import { Toaster } from "$lib/components/ui/sonner/index.js";
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

    let { children } = $props();
</script>

<svelte:head>
    <link rel="icon" href={favicon} />
</svelte:head>

<ModeWatcher />
<Toaster />

<div class="flex min-h-screen flex-col mx-auto max-w-5xl">
    <header class="border-b">
        <div class="flex items-center justify-between p-4">
            <nav class="flex items-center gap-6">
                <a href="/" class={cn("text-lg font-semibold hover:underline")}>
                    zabdoc
                </a>
            </nav>
            <nav class="flex items-center gap-6">
                <Theme />
                <a href="/about" class={buttonVariants({ variant: "ghost" })}>
                    about
                </a>
            </nav>
        </div>
    </header>

    <main class="flex-1 py-8">
        {@render children?.()}
    </main>

    <footer class="border-t">
        <div
            class="p-4 flex items-center justify-between text-sm text-muted-foreground"
        >
            <a
                href="https://github.com/hammadmajid/zabdoc"
                target="_blank"
                rel="noreferrer"
            >
                GitHub
            </a>
            <span
                >Licensed under <a
                    href="https://www.gnu.org/licenses/gpl-3.0.html">GPL v3</a
                ></span
            >
        </div>
    </footer>
</div>
