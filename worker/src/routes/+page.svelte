<script lang="ts">
    import FileText from "@lucide/svelte/icons/file-text";
    import ClipboardList from "@lucide/svelte/icons/clipboard-list";
    import SEO from "$lib/components/seo.svelte";
</script>

<SEO
    title="zabdoc"
    description="Generate assignment and lab task PDFs for SZABIST students."
    canonical="https://zabdoc.xyz/"
    url="https://zabdoc.xyz/"
/>

<div
    class="flex flex-col items-center justify-center text-center h-full px-4 py-12"
>
    <div class="neo-border neo-shadow-lg bg-primary px-8 py-4 mb-8 rotate-[-2deg]">
        <h1 class="text-5xl md:text-7xl font-black uppercase tracking-tight">
            zabdoc
        </h1>
    </div>

    <p class="mt-4 text-lg md:text-xl font-medium max-w-xl">
        web application for
        <span class="bg-secondary px-2 py-0.5 neo-border-sm font-bold"
            >SZABIST*</span
        >
        students to generate assignment and lab tasks in PDF format.
    </p>

    <div class="py-12 flex flex-col sm:flex-row gap-6">
        <a
            href="/assignment"
            class="neo-border neo-shadow bg-secondary text-secondary-foreground px-12 py-6 text-xl font-black uppercase tracking-wide flex items-center gap-3 hover:translate-x-[4px] hover:translate-y-[4px] hover:shadow-none transition-all"
        >
            <FileText class="size-8" />
            Assignment
        </a>
        <a
            href="/lab-task"
            class="neo-border neo-shadow bg-accent text-accent-foreground px-12 py-6 text-xl font-black uppercase tracking-wide flex items-center gap-3 hover:translate-x-[4px] hover:translate-y-[4px] hover:shadow-none transition-all"
        >
            <ClipboardList class="size-8" />
            Lab Task
        </a>
    </div>

    <p class="text-sm font-medium bg-muted neo-border-sm px-4 py-2 max-w-md">
        *this project is <em>not sponsored, affiliated, endorsed, or approved by</em
        > SZABIST.
        <a href="/about#disclaimer" class="underline font-bold hover:text-primary"
            >Learn more.</a
        >
    </p>
</div>
