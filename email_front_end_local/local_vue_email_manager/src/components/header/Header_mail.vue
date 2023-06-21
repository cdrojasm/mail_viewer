<template>
    <div id="container_mail_header" class="drop-shadow-lg">
        <div class="container_image_mail_header rounded-full">
            <img class="header_mail_image" src="@/assets/enron_logo.png" />
        </div>
        <div class="container_textTitle_mail_header">
            <div>
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
        display: flex;
        flex-direction: row;
    }

    .container_image_mail_header {
        width: 20%;
        height: 90%;
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .header_mail_image {
        max-width: 130px;
        max-height: 100px;
    }

    .container_textTitle_mail_header {
        width: 80%;
        height: 90%;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-start;
        min-height: 98px;
    }

    .container_textTitle_mail_header>div {
        height: 100%;
        width: 100%;
    }
}
</style>