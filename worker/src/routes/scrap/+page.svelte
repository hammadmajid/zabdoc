<script lang="ts">
    import { fade, scale } from "svelte/transition";
    import { quintOut } from "svelte/easing";
    import ArrowLeft from "@lucide/svelte/icons/arrow-left";
    import Database from "@lucide/svelte/icons/database";
    import Loader2 from "@lucide/svelte/icons/loader-2";
    import XCircle from "@lucide/svelte/icons/x-circle";
    import Construction from "@lucide/svelte/icons/construction";
    import * as Card from "$lib/components/ui/card";
    import * as Tabs from "$lib/components/ui/tabs";
    import * as Table from "$lib/components/ui/table";
    import Button from "$lib/components/ui/button/button.svelte";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";

    interface AttendanceRecord {
        lecture: string;
        date: string;
        status: string;
    }

    interface CourseAttendance {
        courseName: string;
        instructor: string;
        records: AttendanceRecord[];
    }

    let username = $state("");
    let password = $state("");
    let isLoading = $state(false);
    let isError = $state(false);
    let errorMessage = $state("");
    let courseData = $state<CourseAttendance[]>([]);
    let hasResults = $state(false);

    async function handleScrap() {
        if (!username.trim() || !password.trim()) {
            isError = true;
            errorMessage = "Username and password are required";
            return;
        }

        isLoading = true;
        isError = false;
        errorMessage = "";

        try {
            const apiUrl = window.location.host.includes("localhost")
                ? "http://localhost:8080/scrap"
                : "https://api.zabdoc.xyz/scrap";

            const response = await fetch(apiUrl, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ username, password }),
            });

            if (!response.ok) {
                const errorData: any = await response.json().catch(() => null);
                if (response.status === 400) {
                    throw new Error(errorData?.error || "Invalid username or password");
                } else if (response.status === 500) {
                    throw new Error(errorData?.error || "Server error. Please try again later.");
                } else {
                    throw new Error(errorData?.error || `Error: ${response.status}`);
                }
            }

            const result: any = await response.json();

            console.log("API Response:", result);

            if (!result.success) {
                throw new Error(result.error || "Failed to fetch attendance data");
            }

            // Ensure data is an array with proper structure
            if (Array.isArray(result.data)) {
                courseData = result.data.map((course: any) => ({
                    courseName: course.courseName || "Unknown Course",
                    instructor: course.instructor || "Unknown Instructor",
                    records: Array.isArray(course.records) ? course.records : []
                }));
            } else {
                courseData = [];
            }

            console.log("Processed courseData:", courseData);
            hasResults = true;
        } catch (error) {
            isError = true;
            errorMessage = error instanceof Error ? error.message : String(error);
        } finally {
            isLoading = false;
        }
    }

    function getStatusColor(status: string): string {
        const normalizedStatus = status.toLowerCase().trim();
        if (normalizedStatus === "present" || normalizedStatus === "p") {
            return "text-green-600 bg-green-100 dark:bg-green-950";
        } else if (normalizedStatus === "absent" || normalizedStatus === "a") {
            return "text-red-600 bg-red-100 dark:bg-red-950";
        } else if (normalizedStatus === "late" || normalizedStatus === "l") {
            return "text-yellow-600 bg-yellow-100 dark:bg-yellow-950";
        }
        return "text-muted-foreground bg-muted";
    }
</script>

<div class="min-h-[70vh] flex flex-col">
    <!-- Back to Home Button -->
    <div class="mb-4" in:fade={{ duration: 300 }}>
        <a
            href="/"
            class="neo-border neo-shadow bg-card px-4 py-2 font-bold uppercase text-sm flex items-center gap-2 hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all inline-flex w-fit"
        >
            <ArrowLeft class="size-4" />
            Back to Home
        </a>
    </div>

    {#if !hasResults}
        <!-- Login Form -->
        <div class="flex flex-col items-center justify-center px-4 py-8 flex-1 max-w-md mx-auto w-full">
            <div class="neo-border neo-shadow-lg bg-primary px-8 py-4 mb-8 rotate-[-2deg]">
                <div class="flex items-center gap-3">
                    <Database class="size-6" />
                    <h1 class="text-2xl md:text-3xl font-black uppercase tracking-tight">
                        Scrape Data
                    </h1>
                </div>
            </div>

            <p class="text-base font-medium max-w-md mb-8 text-center">
                Enter your ZabDesk credentials to fetch attendance and marks data.
            </p>

            {#if isError}
                <div
                    class="w-full neo-border neo-shadow-lg bg-destructive/10 p-4 mb-6 border-destructive"
                    in:scale={{ duration: 400, easing: quintOut }}
                >
                    <div class="flex items-center gap-3">
                        <XCircle class="size-5 text-destructive flex-shrink-0" />
                        <p class="text-sm font-medium text-destructive">
                            {errorMessage}
                        </p>
                    </div>
                </div>
            {/if}

            <div class="w-full space-y-4">
                <div class="space-y-2">
                    <Label for="username" class="font-bold uppercase text-sm">Username</Label>
                    <Input
                        id="username"
                        type="text"
                        placeholder="Enter your username"
                        bind:value={username}
                        disabled={isLoading}
                        class="neo-border-sm"
                    />
                </div>

                <div class="space-y-2">
                    <Label for="password" class="font-bold uppercase text-sm">Password</Label>
                    <Input
                        id="password"
                        type="password"
                        placeholder="Enter your password"
                        bind:value={password}
                        disabled={isLoading}
                        class="neo-border-sm"
                        onkeydown={(e) => {
                            if (e.key === "Enter") {
                                handleScrap();
                            }
                        }}
                    />
                </div>

                <Button
                    onclick={handleScrap}
                    disabled={isLoading}
                    class="w-full neo-border neo-shadow bg-primary text-primary-foreground px-6 py-6 text-lg font-black uppercase hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all mt-6"
                >
                    {#if isLoading}
                        <Loader2 class="size-5 mr-2 animate-spin" />
                        Scraping Data...
                    {:else}
                        <Database class="size-5 mr-2" />
                        Scrape Data
                    {/if}
                </Button>
            </div>
        </div>
    {:else}
        <!-- Results Display -->
        <div class="flex flex-col px-4 py-8 max-w-6xl mx-auto w-full">
            <div class="flex items-center justify-between mb-8">
                <div class="neo-border neo-shadow-lg bg-primary px-6 py-3 rotate-[-1deg]">
                    <h1 class="text-2xl md:text-3xl font-black uppercase tracking-tight">
                        Scrapped Data
                    </h1>
                </div>
            </div>

            {#if courseData.length === 0}
                <div class="text-center py-12">
                    <p class="text-lg font-medium text-muted-foreground">
                        No attendance data found.
                    </p>
                </div>
            {:else}
                <div class="grid gap-6 md:grid-cols-2">
                    {#each courseData as course, index}
                        <div
                            in:scale={{ duration: 300, delay: index * 50, easing: quintOut }}
                        >
                            <Card.Root>
                                <Card.Header>
                                    <Card.Title class="text-xl font-black uppercase">
                                        {course.courseName}
                                    </Card.Title>
                                    <Card.Description class="font-medium">
                                        {course.instructor}
                                    </Card.Description>
                                </Card.Header>
                                <Card.Content>
                                    <Tabs.Root value="attendance" class="w-full">
                                        <Tabs.List class="grid w-full grid-cols-2 neo-border-sm">
                                            <Tabs.Trigger value="attendance" class="font-bold uppercase text-sm">
                                                Attendance
                                            </Tabs.Trigger>
                                            <Tabs.Trigger value="marks" class="font-bold uppercase text-sm">
                                                Marks
                                            </Tabs.Trigger>
                                        </Tabs.List>
                                        <Tabs.Content value="attendance" class="mt-4">
                                            {#if !course.records || course.records.length === 0}
                                                <p class="text-sm text-muted-foreground text-center py-4">
                                                    No attendance records available.
                                                </p>
                                            {:else}
                                                <Table.Root>
                                                    <Table.Header>
                                                        <Table.Row>
                                                            <Table.Head>Lecture</Table.Head>
                                                            <Table.Head>Date</Table.Head>
                                                            <Table.Head>Status</Table.Head>
                                                        </Table.Row>
                                                    </Table.Header>
                                                    <Table.Body>
                                                        {#each course.records as record}
                                                            <Table.Row>
                                                                <Table.Cell class="font-medium">
                                                                    {record.lecture}
                                                                </Table.Cell>
                                                                <Table.Cell class="text-muted-foreground">
                                                                    {record.date}
                                                                </Table.Cell>
                                                                <Table.Cell>
                                                                    <span
                                                                        class="inline-block px-2 py-1 rounded text-xs font-bold uppercase {getStatusColor(record.status)}"
                                                                    >
                                                                        {record.status}
                                                                    </span>
                                                                </Table.Cell>
                                                            </Table.Row>
                                                        {/each}
                                                    </Table.Body>
                                                </Table.Root>
                                            {/if}
                                        </Tabs.Content>
                                        <Tabs.Content value="marks" class="mt-4">
                                            <div class="flex flex-col items-center justify-center py-8 text-center space-y-4">
                                                <div class="neo-border-sm bg-muted p-4 inline-block">
                                                    <Construction class="size-12 text-muted-foreground" />
                                                </div>
                                                <p class="text-lg font-black uppercase text-muted-foreground">
                                                    Work in Progress
                                                </p>
                                                <p class="text-sm text-muted-foreground">
                                                    Marks feature is not yet implemented.
                                                </p>
                                            </div>
                                        </Tabs.Content>
                                    </Tabs.Root>
                                </Card.Content>
                            </Card.Root>
                        </div>
                    {/each}
                </div>
            {/if}
        </div>
    {/if}
</div>
