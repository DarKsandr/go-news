{{template "base" .}}

{{define "title"}}404 - Страница не найдена{{end}}

{{define "content"}}
<div class="container-fluid py-5">
    <div class="container py-5 text-center">
        <ol class="breadcrumb justify-content-center mb-5">
            <li class="breadcrumb-item"><a href="/">Главная</a></li>
            <li class="breadcrumb-item active text-dark">404</li>
        </ol>
        <div class="row justify-content-center">
            <div class="col-lg-6">
                <i class="bi bi-exclamation-triangle display-1 text-secondary"></i>
                <h1 class="display-1">404</h1>
                <h1 class="mb-4">Страница не найдена</h1>
                <p class="mb-4">Извините, страница, которую вы искали, отсутствует на нашем сайте! Возможно, вам стоит
                    перейти на нашу
                    главную страницу или воспользоваться поиском?</p>
                <a class="btn link-hover border border-primary rounded-pill py-3 px-5" href="/">Вернуться на главную страницу</a>
            </div>
        </div>
    </div>
</div>
{{end}}