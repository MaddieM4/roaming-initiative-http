<!DOCTYPE html>
<html>
<head>
    <title>{{ title or 'Roaming Initiative Official Site' }}</title>
    <link href="/media/css/main.css" rel="stylesheet" type="text/css">
</head>
<body>
%include templates/header.tpl ip=ip
<div class="container">
    %include templates/sidebar.tpl
    <div id="main-content">
        %include
    </div>
</div>
%include templates/footer.tpl
</body>
</html>
