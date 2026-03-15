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
  <div
    class="prose prose-neutral dark:prose-invert prose-headings:font-black prose-headings:uppercase prose-headings:tracking-tight prose-a:text-primary prose-a:font-bold prose-strong:font-black max-w-none"
  >
    <h1>About</h1>

    <p>
      <strong>zabdoc</strong> is a lightweight web app built for
      <span class="bg-secondary px-1 neo-border-sm font-bold">SZABIST</span>
      students to generate polished assignment cover sheets and lab submissions
      as downloadable PDFs, and to securely fetch attendance and grades from
      ZabDesk using your own credentials.
    </p>

    <h2>How It Works</h2>

    <p>
      zabdoc runs on a <a href="https://go.dev" target="_blank" rel="noreferrer">Go</a>
      backend hosted on <a href="https://www.heroku.com/" target="_blank" rel="noreferrer">Heroku</a>,
      with a <a href="https://svelte.dev" target="_blank" rel="noreferrer">SvelteKit</a>
      frontend deployed to <a href="https://workers.cloudflare.com/" target="_blank" rel="noreferrer">Cloudflare Workers</a>.
      The backend handles PDF generation and ZabDesk scraping, while the frontend
      delivers a fast, responsive UI.
    </p>

    <h3>PDF Generation</h3>
    <p>
      zabdoc generates professional assignment cover sheets and lab submission
      PDFs by rendering a custom HTML template server-side using a headless
      Chrome browser. When you submit the form:
    </p>
    <ol>
      <li>Your submission data is validated and formatted</li>
      <li>A personalized HTML template is generated with your information</li>
      <li>
        The backend uses <a
          href="https://github.com/chromedp/chromedp"
          target="_blank"
          rel="noreferrer">chromedp</a
        > to render the HTML to PDF via headless Chrome
      </li>
      <li>The final PDF is returned to your browser for download</li>
    </ol>
    <p>
      This approach ensures pixel-perfect rendering and consistent output across
      all browsers and devices.
    </p>

    <h3>Data Scraping</h3>
    <p>
      The scraper allows you to fetch your attendance and marks data directly
      from ZabDesk. Your credentials are used only for that request and are not
      stored:
    </p>
    <ol>
      <li>
        You provide your ZabDesk username and password on the scraper page
      </li>
      <li>The credentials are sent securely to our backend API</li>
      <li>
        The backend authenticates with ZabDesk and fetches your enrolled courses
      </li>
      <li>
        For each course, attendance and marks data is extracted and displayed in
        your browser
      </li>
      <li>
        Your credentials are <strong>never stored</strong> - they are only used for
        this request
      </li>
    </ol>

    <h2>Features</h2>

    <ul>
      <li>
        Generate assignment cover sheets and lab submissions in PDF format
      </li>
      <li>
        Scrape attendance and marks data from official ZabDesk using
        user-provided credentials
      </li>
      <li>
        Store basic student details locally for faster repeat form filling
      </li>
    </ul>

    <h2 id="disclaimer">Disclaimer</h2>

    <p>
      <strong>Important:</strong> This project is{" "}
      <em>not sponsored, affiliated, endorsed, or approved by</em> Shaheed Zulfikar
      Ali Bhutto Institute of Science and Technology (SZABIST) Islamabad or any of
      its other campuses. This is an independent, unofficial tool created to assist
      students.
    </p>

    <p>
      <strong>No Warranty:</strong> This program is distributed in the hope that
      it will be useful, but <em>WITHOUT ANY WARRANTY</em>; without even the
      implied warranty of{" "}
      <em>MERCHANTABILITY</em> or <em>FITNESS FOR A PARTICULAR PURPOSE</em>. See
      the
      <a
        href="https://www.gnu.org/licenses/gpl-3.0.html"
        target="_blank"
        rel="noopener noreferrer"
      >
        GNU General Public License</a
      >
      for more details.
    </p>

    <p class="text-sm text-muted-foreground">
      For detailed terms and policies, see our <a
        href="/terms"
        class="text-primary font-bold hover:underline">Terms of Use</a
      >
      and
      <a href="/privacy" class="text-primary font-bold hover:underline"
        >Privacy Policy</a
      >.
    </p>
  </div>

  <!-- FAQ Section -->
  <div>
    <h2 class="text-2xl font-black uppercase tracking-tight mb-6">
      Frequently Asked Questions
    </h2>

    <Accordion.Root type="multiple">
      <Accordion.Item value="item-1">
        <Accordion.Trigger
          >Is this tool officially approved by SZABIST?</Accordion.Trigger
        >
        <Accordion.Content>
          <p class="text-sm">
            No, this is an independent project created to help students. It is
            not affiliated with or endorsed by SZABIST. See the disclaimer above
            for more details.
          </p>
        </Accordion.Content>
      </Accordion.Item>

      <Accordion.Item value="item-2">
        <Accordion.Trigger
          >Can I save my information for future use?</Accordion.Trigger
        >
        <Accordion.Content>
          <p class="text-sm">
            Yes! Your student name and registration number are automatically
            saved locally in your browser for convenience. ZabDesk credentials
            are not stored in local storage.
          </p>
        </Accordion.Content>
      </Accordion.Item>

      <Accordion.Item value="item-2b">
        <Accordion.Trigger
          >Does zabdoc store my ZabDesk username and password?</Accordion.Trigger
        >
        <Accordion.Content>
          <p class="text-sm">
            Credentials are used to log in and fetch your attendance data for
            your request. They are not saved to local storage by zabdoc.
          </p>
        </Accordion.Content>
      </Accordion.Item>

      <Accordion.Item value="item-3">
        <Accordion.Trigger
          >How can I delete my saved information?</Accordion.Trigger
        >
        <Accordion.Content>
          <p class="text-sm">
            Go to the <a
              href="/settings"
              class="text-primary font-bold hover:underline">Settings</a
            >
            page and click on <strong>Delete Local Data</strong>. This will
            clear your saved information from your browser.
          </p>
        </Accordion.Content>
      </Accordion.Item>

      <Accordion.Item value="item-4">
        <Accordion.Trigger
          >My class or course is not listed. What do I do?</Accordion.Trigger
        >
        <Accordion.Content>
          <p class="text-sm mb-2">You have several options:</p>
          <ul class="text-sm list-disc list-inside space-y-1">
            <li>
              Email <a
                href="mailto:hammadmajid@pm.me"
                class="text-primary font-bold hover:underline"
                >hammadmajid@pm.me</a
              > with your request
            </li>
            <li>
              Open an issue on <a
                href="https://github.com/hammadmajid/zabdoc/issues"
                target="_blank"
                rel="noopener noreferrer"
                class="text-primary font-bold hover:underline">GitHub</a
              >
            </li>
            <li>
              Contribute directly via a <a
                href="https://github.com/hammadmajid/zabdoc/tree/main/docs/adding-course.md"
                target="_blank"
                rel="noopener noreferrer"
                class="text-primary font-bold hover:underline">pull request</a
              >
            </li>
          </ul>
          <p class="text-sm mt-2">
            Please provide: class name, course name, course code, and instructor
            name for all courses in the class.
          </p>
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
