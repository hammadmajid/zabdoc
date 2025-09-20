<script lang="ts">
  export let title: string =
    "zabdoc";
  export let description: string =
    "Generate assignment, lab task, and lab project PDFs for SZABIST students.";
  export let canonical: string = "https://zabdoc.xyz/";
  export let url: string = canonical;
  const ogImage = "/og-image.png"; // TODO: OG IMAGE FILE DOES NOT EXIST.
</script>

<svelte:head>
  <title>{title}</title>
  <meta name="description" content={description} />
  <link rel="canonical" href={canonical} />
  <meta property="og:title" content={title} />
  <meta property="og:description" content={description} />
  <meta property="og:type" content="website" />
  <meta property="og:url" content={url} />
  <meta property="og:image" content={ogImage} />
  <meta name="twitter:card" content="summary_large_image" />
  <meta name="twitter:title" content={title} />
  <meta name="twitter:description" content={description} />
  <meta name="twitter:image" content={ogImage} />
</svelte:head>
