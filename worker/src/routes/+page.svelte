<script lang="ts">
    import {buttonVariants} from "$lib/components/ui/button";
</script>

<div class="flex flex-col items-center justify-center text-center py-20">
    <h1 class="text-5xl font-bold tracking-tight">zabdoc</h1>
    <p class="mt-4 text-lg text-muted-foreground max-w-xl">
        Generate professional assignment cover sheets and lab tasks for SZABIST students
    </p>

    <div class="mt-8 flex gap-4">
        <a href="/assignment"
           class={buttonVariants({variant: "default"})}>
            Assignment Cover
        </a>
        <a href="/lab-task"
           class={buttonVariants({variant: "default"})}>
            Lab Task
        </a>
    </div>
</div>
