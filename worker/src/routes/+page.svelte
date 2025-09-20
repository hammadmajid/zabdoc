<script lang="ts">
    import FileText from "@lucide/svelte/icons/file-text";
    import ClipboardList from "@lucide/svelte/icons/clipboard-list";
    import Microscope from "@lucide/svelte/icons/microscope";
    import { buttonVariants } from "$lib/components/ui/button";
    import SEO from "$lib/components/seo.svelte";
    import CardDescription from "$lib/components/ui/card/card-description.svelte";
</script>

<SEO
    title="zabdoc"
    description="Generate assignment, lab task, and lab project PDFs for SZABIST students."
    canonical="https://zabdoc.xyz/"
    url="https://zabdoc.xyz/"
/>

<div class="flex flex-col items-center justify-center text-center h-full px-4">
    <h1 class="text-5xl font-bold tracking-tight">zabdoc</h1>
    <p class="mt-4 text-lg text-muted-foreground max-w-xl">
        web application for
        <span class="border-b border-dotted border-foreground/50">*SZABIST</span>
        students to generate assignment, lab tasks and lab projects in PDF format.
    </p>
    <div class="py-12 grid grid-cols-2 gap-4">
        <a
            href="/assignment"
            class={buttonVariants({
                variant: "outline",
                size: "lg",
                className: "px-12 py-6",
            })}
        >
            <FileText /> Assignment
        </a>
        <a
            href="/lab-task"
            class={buttonVariants({
                variant: "outline",
                size: "lg",
                className: "px-12 py-6",
            })}
        >
            <ClipboardList /> Lab Task
        </a>
        <a
            href="/lab-project"
            class={buttonVariants({
                variant: "outline",
                size: "lg",
                className: "px-12 py-6",
            })}
        >
            <Microscope /> Lab Project
        </a>
    </div>
    <CardDescription
        >*this project is{" "}
        <em>not sponsored, affiliated, endorsed, or approved by</em> SZABIST.
        <a href="/about#disclaimer" class="underline">Learn more.</a></CardDescription
    >
</div>
