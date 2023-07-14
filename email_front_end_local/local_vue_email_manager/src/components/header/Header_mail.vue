<template>
    <div id="container_mail_header" class=" flex flex-row">
        <div class="container_image_mail_header w-1/5 h-5/6	flex flex-col items-center rounded-full">
            <img class="header_mail_image max-h-28	" src="@/assets/enron_logo.png" />
        </div>
        <div class="container_textTitle_mail_header w-full h-5/6 flex flex-row items-center justify-start">
            <div class="h-full w-full">
                <h2 class="text-3xl font-sans leading-7 sm:truncate text-gray-700 sm:text-3xl sm:tracking-tight">Enron Mail
                    Viewer</h2>
            </div>
        </div>
    </div>
</template>
<style scoped>
@media (min-width: 1024px) {
    #container_mail_header {
        min-width: 1000px;
    }

    .header_mail_image {
        max-width: 130px;
    }

    .container_textTitle_mail_header {
        min-height: 98px;
    }
}
</style>