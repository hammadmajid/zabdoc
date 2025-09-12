<script lang="ts">
    import AutoForm from "$lib/components/auto-form.svelte";
    import DueDate from "$lib/components/due-date.svelte";
    import Button from "$lib/components/ui/button/button.svelte";
    import Input from "$lib/components/ui/input/input.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import { Textarea } from "$lib/components/ui/textarea/index.js";

    interface Task {
        id: number;
        content: string;
        files: FileList | null;
    }

    let tasks: Task[] = $state([{ id: 0, content: "", files: null }]);
    let taskCounter = $state(1);

    function addTask() {
        taskCounter++;
        tasks = [...tasks, { id: taskCounter, content: "", files: null }];
    }

    function removeTask(taskId: number) {
        if (tasks.length > 1) {
            tasks = tasks.filter((task) => task.id !== taskId);
        }
    }

    function updateTaskContent(taskId: number, content: string) {
        tasks = tasks.map((task) =>
            task.id === taskId ? { ...task, content } : task,
        );
    }

    function updateTaskFiles(taskId: number, files: FileList | null) {
        tasks = tasks.map((task) =>
            task.id === taskId ? { ...task, files } : task,
        );
    }
</script>

<svelte:head>
    <title>lab task | zabdoc</title>
</svelte:head>

<div class="space-y-12 p-4 md:p-0">
    <div>
        <h1
            class="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
        >
            Lab task
        </h1>
    </div>
    <form
        action="/api/lab-task"
        method="POST"
        enctype="multipart/form-data"
        class="space-y-8"
    >
        <div class="grid md:grid-cols-2 gap-8">
            <AutoForm />

            <Card.Root class="">
                <Card.Header>
                    <Card.Title>Document</Card.Title>
                    <Card.Description
                        >Information about the lab task.
                    </Card.Description>
                </Card.Header>
                <Card.Content class="space-y-4">
                    <DueDate />

                    <Input
                        name="number"
                        type="number"
                        placeholder="Number"
                        required
                    />
                </Card.Content>
            </Card.Root>

            {#each tasks as task, index (task.id)}
                <Card.Root class="relative">
                    <Card.Header>
                        <div class="flex items-center justify-between">
                            <Card.Title>Task {index}</Card.Title>
                            {#if tasks.length > 1}
                                <!-- TODO: confirm this action -->
                                <Button
                                    type="button"
                                    variant="outline"
                                    size="sm"
                                    onclick={() => removeTask(task.id)}
                                >
                                    Remove
                                </Button>
                            {/if}
                        </div>
                        <Card.Description
                            >Task content and files.</Card.Description
                        >
                    </Card.Header>
                    <Card.Content class="space-y-4">
                        <Textarea
                            rows={5}
                            name="task-{index}-content"
                            placeholder="Task content. [Markdown is supported]"
                            bind:value={task.content}
                            oninput={(e) => {
                                const target =
                                    e.target as HTMLTextAreaElement | null;
                                if (target)
                                    updateTaskContent(task.id, target.value);
                            }}
                        />

                        <Input
                            id="task-{index}-files"
                            name="task-{index}-files"
                            type="file"
                            multiple
                            accept="image/*"
                            onchange={(e) => {
                                const target =
                                    e.target as HTMLInputElement | null;
                                updateTaskFiles(task.id, target?.files ?? null);
                            }}
                        />
                    </Card.Content>
                </Card.Root>
            {/each}

            {#if tasks.length < 15}
                <Card.Root class="md:min-h-[250px]">
                    <Card.Content class="h-full">
                        <Button
                            class="w-full h-full"
                            type="button"
                            variant="outline"
                            onclick={addTask}
                        >
                            Add Task
                        </Button>
                    </Card.Content>
                </Card.Root>
            {/if}
        </div>

        <Card.Root class="">
            <Card.Header>
                <Card.Description
                    >By clicking the button below you accept terms and privacy/
                </Card.Description>
            </Card.Header>
            <Card.CardContent>
                <Button type="submit" class="w-full">Generate cover</Button>
            </Card.CardContent>
        </Card.Root>
    </form>
</div>
