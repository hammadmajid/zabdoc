<script lang="ts">
    import CardDescription from "$lib/components/ui/card/card-description.svelte";
    import * as Select from "$lib/components/ui/select/index.js";
    import Input from "$lib/components/ui/input/input.svelte";
    import * as Card from "$lib/components/ui/card/index.js";

    type CourseInfo = {
        instructor: string;
        code: string;
    };

    type ClassData = Record<string, CourseInfo>;

    type DataStructure = {
        "BsCS-4C": ClassData;
        "BsCS-5A": ClassData;
    };

    const data: DataStructure = {
        "BsCS-4C": {
            "Database Systems": {
                instructor: "Dr. Shahzad Latif",
                code: "CSC 2203",
            },
            "Finite Automata Theory and Formal Languages": {
                instructor: "Dr. Shahzad Latif",
                code: "CSC 2204",
            },
            "Linear Algebra": {
                instructor: "Dr. Shahzad Latif",
                code: "CSC 2206",
            },
            "Design and Analysis of Algorithms": {
                instructor: "Dr. Shahzad Latif",
                code: "CSC 3202",
            },
        },
        "BsCS-5A": {
            "Natural Language Processing": {
                instructor: "Saad Irfan khan",
                code: "CSC 4207",
            },
            "Technical and Business Writing": {
                instructor: "Muhammad Qasim",
                code: "ENG 3201",
            },
            "Computer Networks and Data Communications": {
                instructor: "Rana Faisal Hayat",
                code: "CSC 4203",
            },
            "Artificial Intelligence": {
                instructor: "Muhammad Qasim",
                code: "CSC 4204",
            },
            "Web Technologies-I": {
                instructor: "Muhammad Qasim",
                code: "CSC 4205",
            },
            "Android Application Development": {
                instructor: "Zubair Ahmed Chatta",
                code: "CSC 4206",
            },
            "Lab: Computer Networks and Data Communications": {
                instructor: "Muhammad Qasim",
                code: "CSC 4208",
            },
            "Lab: Artificial Intelligence": {
                instructor: "Muhammad Qasim",
                code: "CSC 4209",
            },
        },
    };

    let selectedClass = $state("" as keyof DataStructure | "");
    let selectedCourse = $state("");

    const classes = $derived(
        Object.keys(data).map((className) => ({
            value: className,
            label: className,
        })),
    );

    const courses = $derived(
        selectedClass && selectedClass in data
            ? Object.keys(data[selectedClass]).map((courseName) => ({
                  value: courseName,
                  label: courseName,
              }))
            : [],
    );

    const courseDetails = $derived(
        selectedClass &&
            selectedCourse &&
            selectedClass in data &&
            selectedCourse in data[selectedClass]
            ? data[selectedClass][selectedCourse]
            : null,
    );

    const classTriggerContent = $derived(
        classes.find((c) => c.value === selectedClass)?.label ?? "Select class",
    );

    const courseTriggerContent = $derived(
        courses.find((c) => c.value === selectedCourse)?.label ??
            "Select course",
    );

    // Reset course selection when class changes
    $effect(() => {
        if (selectedClass) {
            selectedCourse = "";
        }
    });
</script>

<Card.Root class="">
    <Card.Header>
        <Card.Title>Student</Card.Title>
        <CardDescription
            >This information will be cached for you convenience.
        </CardDescription>
    </Card.Header>
    <Card.Content class="space-y-4">
        <Input name="studentName" type="text" placeholder="Full name" />
        <Input name="regNo" type="number" placeholder="Registration no" />

        <!-- Class Selection -->
        <Select.Root
            type="single"
            name="class"
            required
            bind:value={selectedClass}
        >
            <Select.Trigger class="w-full">
                {classTriggerContent}
            </Select.Trigger>
            <Select.Content>
                <Select.Group>
                    <Select.Label>Classes</Select.Label>
                    {#each classes as classItem (classItem.value)}
                        <Select.Item
                            value={classItem.value}
                            label={classItem.label}
                        >
                            {classItem.label}
                        </Select.Item>
                    {/each}
                </Select.Group>
            </Select.Content>
        </Select.Root>
    </Card.Content>
</Card.Root>

<Card.Root class="">
    <Card.Header>
        <Card.Title>Course</Card.Title>
        <Card.Description
            >Teacher and course code will be auto filled for you.
        </Card.Description>
    </Card.Header>
    <Card.Content class="space-y-4">
        <!-- Course Selection -->
        <Select.Root
            type="single"
            name="course"
            required
            bind:value={selectedCourse}
            disabled={!selectedClass}
        >
            <Select.Trigger class="w-full">
                {courseTriggerContent}
            </Select.Trigger>
            <Select.Content>
                <Select.Group>
                    <Select.Label>Courses</Select.Label>
                    {#each courses as course (course.value)}
                        <Select.Item value={course.value} label={course.label}>
                            {course.label}
                        </Select.Item>
                    {/each}
                </Select.Group>
            </Select.Content>
        </Select.Root>

        <!-- Auto-filled fields -->
        {#if courseDetails}
            <Input
                name="instructor"
                type="text"
                placeholder="Instructor"
                value={courseDetails.instructor}
                readonly
            />
            <Input
                name="courseCode"
                type="text"
                placeholder="Course Code"
                value={courseDetails.code}
                readonly
            />
        {/if}
    </Card.Content>
</Card.Root>
