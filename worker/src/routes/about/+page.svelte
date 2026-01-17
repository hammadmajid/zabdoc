<script lang="ts">
    import SEO from "$lib/components/seo.svelte";
    import * as Accordion from "$lib/components/ui/accordion/index.js";
</script>

<SEO
    title="about | zabdoc"
    description="Learn about zabdoc, an independent tool for SZABIST students to generate assignment and lab PDFs. Not affiliated with SZABIST."
    canonical="https://zabdoc.xyz/about"
    url="https://zabdoc.xyz/about"
/>

<div class="mx-auto space-y-12">
    <div class="prose prose-neutral dark:prose-invert prose-headings:font-black prose-headings:uppercase prose-headings:tracking-tight prose-a:text-primary prose-a:font-bold prose-strong:font-black">
        <h1>About</h1>

        <p>
            <strong>zabdoc</strong> is a web application for
            <span class="bg-secondary px-1 neo-border-sm font-bold">SZABIST</span>
            students to generate assignment cover sheets and lab tasks in PDF format.
        </p>

        <p>
            Written in{" "}
            <a href="https://go.dev" target="_blank" rel="noreferrer">Go</a> and
            <a href="https://svelte.dev" target="_blank" rel="noreferrer">Svelte</a>,
            it provides a simple, responsive web form and produces
            institution-branded PDFs for assignment cover and lab tasks.
        </p>

        <h2 id="disclaimer">Disclaimer</h2>

        <p>
            <strong>Important:</strong> This project is{" "}
            <em>not sponsored, affiliated, endorsed, or approved by</em> Shaheed Zulfikar
            Ali Bhutto Institute of Science and Technology (SZABIST) Islamabad or any
            of its other campuses. This is an independent, unofficial tool created to
            assist students.
        </p>

        <p>
            <strong>No Warranty:</strong> This program is distributed in the hope
            that it will be useful, but <em>WITHOUT ANY WARRANTY</em>; without even
            the implied warranty of{" "}
            <em>MERCHANTABILITY</em> or <em>FITNESS FOR A PARTICULAR PURPOSE</em>.
            See the
            <a
                href="https://www.gnu.org/licenses/gpl-3.0.html"
                target="_blank"
                rel="noopener noreferrer"
            >
                GNU General Public License</a
            >
            for more details.
        </p>

        <h2>Terms of Use</h2>

        <ul>
            <li>This tool is provided for educational and convenience purposes only</li>
            <li>Users are responsible for ensuring the accuracy of information entered</li>
            <li>The generated PDFs should comply with your institution's submission requirements</li>
            <li>Use of this service is at your own risk and discretion</li>
        </ul>

        <h2>Privacy</h2>

        <ul>
            <li><strong>Data Collection:</strong> We only collect the information you provide to generate PDFs</li>
            <li><strong>Data Sharing:</strong> We do not share your personal information with third parties</li>
            <li><strong>Local Storage:</strong> Your name and registration number are saved locally for convenience</li>
        </ul>
    </div>

    <!-- FAQ Section -->
    <div>
        <h2 class="text-2xl font-black uppercase tracking-tight mb-6">Frequently Asked Questions</h2>

        <Accordion.Root type="multiple">
            <Accordion.Item value="item-1">
                <Accordion.Trigger>Is this tool officially approved by SZABIST?</Accordion.Trigger>
                <Accordion.Content>
                    <p class="text-sm">
                        No, this is an independent project created to help students. It is not
                        affiliated with or endorsed by SZABIST. See the disclaimer above for more details.
                    </p>
                </Accordion.Content>
            </Accordion.Item>

            <Accordion.Item value="item-2">
                <Accordion.Trigger>Can I save my information for future use?</Accordion.Trigger>
                <Accordion.Content>
                    <p class="text-sm">
                        Yes! Your student name and registration number are automatically saved
                        locally in your browser for convenience. This data never leaves your device.
                    </p>
                </Accordion.Content>
            </Accordion.Item>

            <Accordion.Item value="item-3">
                <Accordion.Trigger>How can I delete my saved information?</Accordion.Trigger>
                <Accordion.Content>
                    <p class="text-sm">
                        Go to the <a href="/settings" class="text-primary font-bold hover:underline">Settings</a> page
                        and click on <strong>Delete Local Data</strong>. This will clear your saved information from your browser.
                    </p>
                </Accordion.Content>
            </Accordion.Item>

            <Accordion.Item value="item-4">
                <Accordion.Trigger>My class or course is not listed. What do I do?</Accordion.Trigger>
                <Accordion.Content>
                    <p class="text-sm mb-2">You have several options:</p>
                    <ul class="text-sm list-disc list-inside space-y-1">
                        <li>Email <a href="mailto:hammadmajid@pm.me" class="text-primary font-bold hover:underline">hammadmajid@pm.me</a> with your request</li>
                        <li>Open an issue on <a href="https://github.com/hammadmajid/zabdoc/issues" target="_blank" rel="noopener noreferrer" class="text-primary font-bold hover:underline">GitHub</a></li>
                        <li>Contribute directly via a <a href="https://github.com/hammadmajid/zabdoc/tree/main/docs/adding-course.md" target="_blank" rel="noopener noreferrer" class="text-primary font-bold hover:underline">pull request</a></li>
                    </ul>
                    <p class="text-sm mt-2">Please provide: class name, course name, course code, and instructor name for all courses in the class.</p>
                </Accordion.Content>
            </Accordion.Item>

            <Accordion.Item value="item-5">
                <Accordion.Trigger>Can I modify the generated PDF?</Accordion.Trigger>
                <Accordion.Content>
                    <p class="text-sm">
                        The generated PDFs are in final format. If you need changes,
                        regenerate the document with corrected information.
                    </p>
                </Accordion.Content>
            </Accordion.Item>

            <Accordion.Item value="item-6">
                <Accordion.Trigger>PDF not downloading?</Accordion.Trigger>
                <Accordion.Content>
                    <ul class="text-sm list-disc list-inside space-y-1">
                        <li>Check that all required fields are filled correctly</li>
                        <li>Ensure your browser allows downloads from this site</li>
                        <li>Try refreshing the page and submitting again</li>
                        <li>Disable browser extensions that might block downloads</li>
                        <li>Clear your browser cache if issues persist</li>
                    </ul>
                </Accordion.Content>
            </Accordion.Item>
        </Accordion.Root>
    </div>

    <!-- Getting Support Section -->
    <div class="neo-border neo-shadow bg-muted p-6">
        <h2 class="text-xl font-black uppercase tracking-tight mb-4">Need Help?</h2>
        <p class="font-medium mb-4">
            If you encounter technical issues or have suggestions:
        </p>
        <div class="flex flex-wrap gap-3">
            <a
                href="https://github.com/hammadmajid/zabdoc/issues"
                target="_blank"
                rel="noopener noreferrer"
                class="neo-border-sm neo-shadow-sm bg-card px-4 py-2 font-bold uppercase text-sm hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all"
            >
                Report a Bug
            </a>
            <a
                href="mailto:hammadmajid@pm.me"
                class="neo-border-sm neo-shadow-sm bg-primary px-4 py-2 font-bold uppercase text-sm hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all"
            >
                Contact
            </a>
        </div>
    </div>
</div>
