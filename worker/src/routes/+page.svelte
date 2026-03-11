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

<div class="min-h-[70vh] flex flex-col">
    <!-- Hero Section -->
    <div class="text-center mb-4" in:fade={{ duration: 300 }}>
        <div class="neo-border neo-shadow-lg bg-primary px-8 py-4 inline-block rotate-[-2deg] mb-6">
            <h1 class="text-4xl md:text-6xl font-black uppercase tracking-tight">
                zabdoc
            </h1>
        </div>
        <p class="text-lg font-medium max-w-xl mx-auto">
            utility tools for
            <span class="bg-secondary px-2 py-0.5 neo-border-sm font-bold">*SZABIST</span>
            students.
        </p>
    </div>

    <!-- Choice Section -->
    <div class="flex flex-col items-center justify-center text-center px-4 py-8 flex-1">
        <div class="neo-border neo-shadow-lg bg-primary px-8 py-4 mb-8 rotate-[-2deg]">
            <h2 class="text-4xl md:text-5xl font-black uppercase tracking-tight">
                What do you need?
            </h2>
        </div>

        <p class="text-lg font-medium max-w-md mb-8">
            Choose an option to get started.
        </p>

        <div class="flex flex-col sm:flex-row gap-6">
            <a
                href="/scrap"
                class="neo-border neo-shadow bg-accent text-secondary-foreground px-10 py-6 text-xl font-black uppercase tracking-wide flex items-center gap-3 hover:translate-x-[4px] hover:translate-y-[4px] hover:shadow-none transition-all cursor-pointer"
            >
                <Database class="size-8" />
                Scrap Data
            </a>
            <a
                href="/document"
                class="neo-border neo-shadow bg-secondary text-accent-foreground px-10 py-6 text-xl font-black uppercase tracking-wide flex items-center gap-3 hover:translate-x-[4px] hover:translate-y-[4px] hover:shadow-none transition-all cursor-pointer"
            >
                <FileText class="size-8" />
                Generate Document
            </a>
        </div>
    </div>

    <!-- Disclaimer -->
    <div class="text-center mt-4" in:fade={{ duration: 300, delay: 200 }}>
        <p class="text-sm font-medium bg-muted neo-border-sm px-4 py-2 max-w-md inline-block">
            *this project is <em>not sponsored, affiliated, endorsed, or approved by</em> SZABIST.
            <a href="/about#disclaimer" class="underline font-bold hover:text-primary">Learn more.</a>
        </p>
    </div>
</div>
