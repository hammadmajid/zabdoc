<script lang="ts">
    import * as Card from "$lib/components/ui/card/index";
    import FileText from "@lucide/svelte/icons/file-text";
    import ClipboardList from "@lucide/svelte/icons/clipboard-list";
    import Microscope from "@lucide/svelte/icons/microscope";
    import { buttonVariants } from "$lib/components/ui/button";
    import CardDescription from "$lib/components/ui/card/card-description.svelte";
</script>

<div class="flex flex-col items-center justify-center text-center h-full px-4">
    <h1 class="text-5xl font-bold tracking-tight">zabdoc</h1>
    <p class="mt-4 text-lg text-muted-foreground max-w-xl">
        web application for
        <span class="border-b border-dotted border-foreground/50">SZABIST</span>
        students to generate assignment, lab tasks and lab projects in PDF format.
    </p>
    <div class="py-12 grid grid-cols-2 gap-4">
        <a
            href="/assignment"
            class={buttonVariants({ variant: "outline", size: "lg" })}
        >
            <FileText /> Assignment
        </a>
        <a
            href="/lab-task"
            class={buttonVariants({ variant: "outline", size: "lg" })}
        >
            <ClipboardList /> Lab Task
        </a>
        <a
            href="/lab-project"
            class={buttonVariants({ variant: "outline", size: "lg" })}
        >
            <Microscope /> Lab Project [WIP]
        </a>
    </div>
    <div>
        <CardDescription>*WIP: Work in Progress</CardDescription>
    </div>
</div>
