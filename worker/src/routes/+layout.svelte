<script lang="ts">
    import { onNavigate } from "$app/navigation";
    import { page } from "$app/stores";
    import favicon from "$lib/assets/favicon.svg";
    import Settings from "@lucide/svelte/icons/settings";
    import Github from "@lucide/svelte/icons/github";
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

<div class="flex min-h-screen flex-col mx-auto max-w-5xl px-4">
    <header class="py-6">
        <div class="flex items-center justify-between">
            <nav class="flex items-center gap-4">
                <a
                    href="/"
                    class="neo-border-sm neo-shadow-sm bg-primary px-4 py-2 text-xl font-black uppercase tracking-tight hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all"
                >
                    zabdoc
                </a>

            </nav>
            <nav class="flex items-center gap-2">
                <a
                    href="/about"
                    class={cn(
                        "px-3 py-1 text-sm font-bold uppercase tracking-wide transition-colors hover:text-primary",
                        $page.url.pathname === "/about" && "underline"
                    )}
                >
                    about
                </a>
                <a
                    href="/settings"
                    class={buttonVariants({ variant: "ghost", size: "icon" })}
                >
                    <Settings class="size-5" />
                </a>
            </nav>
        </div>
    </header>

    <main class="flex-1 py-6">
        {@render children?.()}
    </main>

    <footer class="py-6 mt-4">
        <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
                <a
                    href="https://github.com/hammadmajid/zabdoc"
                    target="_blank"
                    rel="noreferrer"
                    class="neo-border-sm bg-card p-2 hover:bg-muted transition-colors"
                >
                    <Github class="size-5" />
                </a>
                <span class="text-sm font-medium text-muted-foreground">
                    Open Source
                </span>
            </div>
            <div class="text-sm font-medium">
                <span class="text-muted-foreground">GPL v3</span>
                <span class="mx-2 text-muted-foreground">•</span>
                <span class="bg-muted px-2 py-0.5 neo-border-sm text-xs font-bold">
                    2026
                </span>
            </div>
        </div>
    </footer>
</div>
