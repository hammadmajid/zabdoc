<script lang="ts">
    import * as Card from "$lib/components/ui/card/index";
    import { FileText, ClipboardList } from "@lucide/svelte/icons";
</script>

<div class="flex flex-col items-center justify-center text-center h-full px-4">
    <h1 class="text-5xl font-bold tracking-tight">zabdoc</h1>
    <p class="mt-4 text-lg text-muted-foreground max-w-xl">
        Generate professional assignment cover sheets and lab tasks for SZABIST
        students
    </p>

    <div class="mt-12 grid grid-cols-1 md:grid-cols-2 gap-6 w-full max-w-2xl">
        <a href="/assignment" class="group">
            <Card.Root class="h-full transition-all duration-200 hover:shadow-lg hover:scale-105 cursor-pointer border-2 hover:border-primary/20">
                <Card.Header class="text-center pb-2">
                    <div class="mx-auto mb-4 p-3 bg-primary/10 rounded-full w-fit group-hover:bg-primary/20 transition-colors">
                        <FileText class="h-8 w-8 text-primary" />
                    </div>
                    <Card.Title class="text-xl">Assignment Cover</Card.Title>
                </Card.Header>
                <Card.Content class="text-center">
                    <Card.Description class="text-base leading-relaxed">
                        Create professional cover sheets for your assignments with student details, course information, and due dates.
                    </Card.Description>
                </Card.Content>
            </Card.Root>
        </a>

        <a href="/lab-task" class="group">
            <Card.Root class="h-full transition-all duration-200 hover:shadow-lg hover:scale-105 cursor-pointer border-2 hover:border-primary/20">
                <Card.Header class="text-center pb-2">
                    <div class="mx-auto mb-4 p-3 bg-primary/10 rounded-full w-fit group-hover:bg-primary/20 transition-colors">
                        <ClipboardList class="h-8 w-8 text-primary" />
                    </div>
                    <Card.Title class="text-xl">Lab Task</Card.Title>
                </Card.Header>
                <Card.Content class="text-center">
                    <Card.Description class="text-base leading-relaxed">
                        Generate comprehensive lab task covers with multiple tasks, file attachments, and detailed content sections.
                    </Card.Description>
                </Card.Content>
            </Card.Root>
        </a>
    </div>
</div>
