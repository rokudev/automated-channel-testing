<!DOCTYPE html>
<html class="" lang="en">
<head prefix="og: http://ogp.me/ns#">
<meta charset="utf-8">
<meta content="IE=edge" http-equiv="X-UA-Compatible">
<meta content="object" property="og:type">
<meta content="GitLab" property="og:site_name">
<meta content="RobotLibrary/Tests/test_7_EpisodePicker_utils.robot · develop · developer_web_tools / roku-automated-channel-testing" property="og:title">
<meta content="Roku Automated Channel Testing: Selenium Driver + Robot Framework + Samples" property="og:description">
<meta content="/uploads/-/system/project/avatar/3199/ic_selenium_driver_512.png" property="og:image">
<meta content="64" property="og:image:width">
<meta content="64" property="og:image:height">
<meta content="https://gitlab.eng.roku.com/developer_web_tools/roku-automated-channel-testing/blob/develop/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot" property="og:url">
<meta content="summary" property="twitter:card">
<meta content="RobotLibrary/Tests/test_7_EpisodePicker_utils.robot · develop · developer_web_tools / roku-automated-channel-testing" property="twitter:title">
<meta content="Roku Automated Channel Testing: Selenium Driver + Robot Framework + Samples" property="twitter:description">
<meta content="/uploads/-/system/project/avatar/3199/ic_selenium_driver_512.png" property="twitter:image">

<title>RobotLibrary/Tests/test_7_EpisodePicker_utils.robot · develop · developer_web_tools / roku-automated-channel-testing · GitLab</title>
<meta content="Roku Automated Channel Testing: Selenium Driver + Robot Framework + Samples" name="description">
<link rel="shortcut icon" type="image/png" href="/assets/favicon-7901bd695fb93edb07975966062049829afb56cf11511236e61bcf425070e36e.png" id="favicon" data-original-href="/assets/favicon-7901bd695fb93edb07975966062049829afb56cf11511236e61bcf425070e36e.png" />
<link rel="stylesheet" media="all" href="/assets/application-318ee33e5d14035b04832fa07c492cdf57788adda50bb5219ef75b735cbf00e2.css" />
<link rel="stylesheet" media="print" href="/assets/print-74c3df10dad473d66660c828e3aa54ca3bfeac6d8bb708643331403fe7211e60.css" />



<link rel="stylesheet" media="all" href="/assets/highlight/themes/white-a165d47ce52cf24c29686366976ae691bd9addb9641a6abeb3ba6d1823b89aa8.css" />
<script>
//<![CDATA[
window.gon={};gon.api_version="v4";gon.default_avatar_url="https://gitlab.eng.roku.com/assets/no_avatar-849f9c04a3a0d0cea2424ae97b27447dc64a7dbfae83c036c45b403392f0e8ba.png";gon.max_file_size=10;gon.asset_host=null;gon.webpack_public_path="/assets/webpack/";gon.relative_url_root="";gon.shortcuts_path="/help/shortcuts";gon.user_color_scheme="white";gon.gitlab_url="https://gitlab.eng.roku.com";gon.revision="e2d0c51ffc5";gon.gitlab_logo="/assets/gitlab_logo-7ae504fe4f68fdebb3c2034e36621930cd36ea87924c11ff65dbcb8ed50dca58.png";gon.sprite_icons="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg";gon.sprite_file_icons="/assets/file_icons-7262fc6897e02f1ceaf8de43dc33afa5e4f9a2067f4f68ef77dcc87946575e9e.svg";gon.emoji_sprites_css_path="/assets/emoji_sprites-289eccffb1183c188b630297431be837765d9ff4aed6130cf738586fb307c170.css";gon.test_env=false;gon.suggested_label_colors=["#0033CC","#428BCA","#44AD8E","#A8D695","#5CB85C","#69D100","#004E00","#34495E","#7F8C8D","#A295D6","#5843AD","#8E44AD","#FFECDB","#AD4363","#D10069","#CC0033","#FF0000","#D9534F","#D1D100","#F0AD4E","#AD8D43"];gon.first_day_of_week=0;gon.ee=true;gon.current_user_id=706;gon.current_username="jduval";gon.current_user_fullname="Jonathan Duval";gon.current_user_avatar_url="https://secure.gravatar.com/avatar/b5109351552b7c8a8722d7d971fc4e8f?s=80\u0026d=identicon";
//]]>
</script>


<script src="/assets/webpack/runtime.901bebd8.bundle.js" defer="defer"></script>
<script src="/assets/webpack/main.e5eb6273.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons~pages.admin.clusters~pages.admin.clusters.destroy~pages.admin.clusters.edit~pages.admin.clus~a2ef139c.04bf4fba.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons~pages.groups.epics.index~pages.groups.epics.show~pages.groups.milestones.edit~pages.groups.m~14875979.a05c0e92.chunk.js" defer="defer"></script>
<script src="/assets/webpack/pages.projects.blob.show.ee03c3bb.chunk.js" defer="defer"></script>
<script>
  window.uploads_path = "/developer_web_tools/roku-automated-channel-testing/uploads";
</script>

<meta name="csrf-param" content="authenticity_token" />
<meta name="csrf-token" content="aJqcSSQBA0/3b5QRbQxeOHkI3JczutHNHHm3QCeyXg7mmnS2cb0ncKfe+oxldH+jPwPGgl4zL2f+WYTYWofqtQ==" />
<meta content="origin-when-cross-origin" name="referrer">
<meta content="width=device-width, initial-scale=1, maximum-scale=1" name="viewport">
<meta content="#474D57" name="theme-color">
<link rel="apple-touch-icon" type="image/x-icon" href="/assets/touch-icon-iphone-5a9cee0e8a51212e70b90c87c12f382c428870c0ff67d1eb034d884b78d2dae7.png" />
<link rel="apple-touch-icon" type="image/x-icon" href="/assets/touch-icon-ipad-a6eec6aeb9da138e507593b464fdac213047e49d3093fc30e90d9a995df83ba3.png" sizes="76x76" />
<link rel="apple-touch-icon" type="image/x-icon" href="/assets/touch-icon-iphone-retina-72e2aadf86513a56e050e7f0f2355deaa19cc17ed97bbe5147847f2748e5a3e3.png" sizes="120x120" />
<link rel="apple-touch-icon" type="image/x-icon" href="/assets/touch-icon-ipad-retina-8ebe416f5313483d9c1bc772b5bbe03ecad52a54eba443e5215a22caed2a16a2.png" sizes="152x152" />
<link color="rgb(226, 67, 41)" href="/assets/logo-d36b5212042cebc89b96df4bf6ac24e43db316143e89926c0db839ff694d2de4.svg" rel="mask-icon">
<meta content="/assets/msapplication-tile-1196ec67452f618d39cdd85e2e3a542f76574c071051ae7effbfde01710eb17d.png" name="msapplication-TileImage">
<meta content="#30353E" name="msapplication-TileColor">




</head>

<body class="ui-indigo  gl-browser-chrome gl-platform-mac" data-find-file="/developer_web_tools/roku-automated-channel-testing/find_file/develop" data-group="" data-page="projects:blob:show" data-project="roku-automated-channel-testing">

<script>
  gl = window.gl || {};
  gl.client = {"isChrome":true,"isMac":true};
</script>



<header class="navbar navbar-gitlab qa-navbar navbar-expand-sm js-navbar">
<a class="sr-only gl-accessibility" href="#content-body" tabindex="1">Skip to content</a>
<div class="container-fluid">
<div class="header-content">
<div class="title-container">
<h1 class="title">
<a title="Dashboard" id="logo" href="/"><svg width="24" height="24" class="tanuki-logo" viewBox="0 0 36 36">
  <path class="tanuki-shape tanuki-left-ear" fill="#e24329" d="M2 14l9.38 9v-9l-4-12.28c-.205-.632-1.176-.632-1.38 0z"/>
  <path class="tanuki-shape tanuki-right-ear" fill="#e24329" d="M34 14l-9.38 9v-9l4-12.28c.205-.632 1.176-.632 1.38 0z"/>
  <path class="tanuki-shape tanuki-nose" fill="#e24329" d="M18,34.38 3,14 33,14 Z"/>
  <path class="tanuki-shape tanuki-left-eye" fill="#fc6d26" d="M18,34.38 11.38,14 2,14 6,25Z"/>
  <path class="tanuki-shape tanuki-right-eye" fill="#fc6d26" d="M18,34.38 24.62,14 34,14 30,25Z"/>
  <path class="tanuki-shape tanuki-left-cheek" fill="#fca326" d="M2 14L.1 20.16c-.18.565 0 1.2.5 1.56l17.42 12.66z"/>
  <path class="tanuki-shape tanuki-right-cheek" fill="#fca326" d="M34 14l1.9 6.16c.18.565 0 1.2-.5 1.56L18 34.38z"/>
</svg>

<span class="logo-text d-none d-lg-block prepend-left-8">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 617 169"><path d="M315.26 2.97h-21.8l.1 162.5h88.3v-20.1h-66.5l-.1-142.4M465.89 136.95c-5.5 5.7-14.6 11.4-27 11.4-16.6 0-23.3-8.2-23.3-18.9 0-16.1 11.2-23.8 35-23.8 4.5 0 11.7.5 15.4 1.2v30.1h-.1m-22.6-98.5c-17.6 0-33.8 6.2-46.4 16.7l7.7 13.4c8.9-5.2 19.8-10.4 35.5-10.4 17.9 0 25.8 9.2 25.8 24.6v7.9c-3.5-.7-10.7-1.2-15.1-1.2-38.2 0-57.6 13.4-57.6 41.4 0 25.1 15.4 37.7 38.7 37.7 15.7 0 30.8-7.2 36-18.9l4 15.9h15.4v-83.2c-.1-26.3-11.5-43.9-44-43.9M557.63 149.1c-8.2 0-15.4-1-20.8-3.5V70.5c7.4-6.2 16.6-10.7 28.3-10.7 21.1 0 29.2 14.9 29.2 39 0 34.2-13.1 50.3-36.7 50.3m9.2-110.6c-19.5 0-30 13.3-30 13.3v-21l-.1-27.8h-21.3l.1 158.5c10.7 4.5 25.3 6.9 41.2 6.9 40.7 0 60.3-26 60.3-70.9-.1-35.5-18.2-59-50.2-59M77.9 20.6c19.3 0 31.8 6.4 39.9 12.9l9.4-16.3C114.5 6 97.3 0 78.9 0 32.5 0 0 28.3 0 85.4c0 59.8 35.1 83.1 75.2 83.1 20.1 0 37.2-4.7 48.4-9.4l-.5-63.9V75.1H63.6v20.1h38l.5 48.5c-5 2.5-13.6 4.5-25.3 4.5-32.2 0-53.8-20.3-53.8-63-.1-43.5 22.2-64.6 54.9-64.6M231.43 2.95h-21.3l.1 27.3v94.3c0 26.3 11.4 43.9 43.9 43.9 4.5 0 8.9-.4 13.1-1.2v-19.1c-3.1.5-6.4.7-9.9.7-17.9 0-25.8-9.2-25.8-24.6v-65h35.7v-17.8h-35.7l-.1-38.5M155.96 165.47h21.3v-124h-21.3v124M155.96 24.37h21.3V3.07h-21.3v21.3"/></svg>

</span>
</a></h1>
<ul class="list-unstyled navbar-sub-nav">
<li id="nav-projects-dropdown" class="home dropdown header-projects qa-projects-dropdown" data-track-label="projects_dropdown" data-track-event="click_dropdown"><button class="btn" data-toggle="dropdown" type="button">
Projects
<svg class="caret-down"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#angle-down"></use></svg>
</button>
<div class="dropdown-menu frequent-items-dropdown-menu">
<div class="frequent-items-dropdown-container">
<div class="frequent-items-dropdown-sidebar qa-projects-dropdown-sidebar">
<ul>
<li class=""><a class="qa-your-projects-link" href="/dashboard/projects">Your projects
</a></li><li class=""><a href="/dashboard/projects/starred">Starred projects
</a></li><li class=""><a href="/explore">Explore projects
</a></li></ul>
</div>
<div class="frequent-items-dropdown-content">
<div data-project-avatar-url="/uploads/-/system/project/avatar/3199/ic_selenium_driver_512.png" data-project-id="3199" data-project-name="roku-automated-channel-testing" data-project-namespace="developer_web_tools / roku-automated-channel-testing" data-project-web-url="/developer_web_tools/roku-automated-channel-testing" data-user-name="jduval" id="js-projects-dropdown"></div>
</div>
</div>

</div>
</li><li id="nav-groups-dropdown" class="home dropdown header-groups qa-groups-dropdown" data-track-label="groups_dropdown" data-track-event="click_dropdown"><button class="btn" data-toggle="dropdown" type="button">
Groups
<svg class="caret-down"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#angle-down"></use></svg>
</button>
<div class="dropdown-menu frequent-items-dropdown-menu">
<div class="frequent-items-dropdown-container">
<div class="frequent-items-dropdown-sidebar qa-groups-dropdown-sidebar">
<ul>
<li class=""><a class="qa-your-groups-link" href="/dashboard/groups">Your groups
</a></li><li class=""><a href="/explore/groups">Explore groups
</a></li></ul>
</div>
<div class="frequent-items-dropdown-content">
<div data-user-name="jduval" id="js-groups-dropdown"></div>
</div>
</div>

</div>
</li><li class="d-none d-xl-block"><a class="dashboard-shortcuts-activity" title="Activity" href="/dashboard/activity">Activity
</a></li><li class="d-none d-xl-block"><a class="dashboard-shortcuts-milestones" title="Milestones" href="/dashboard/milestones">Milestones
</a></li><li class="d-none d-xl-block"><a class="dashboard-shortcuts-snippets qa-snippets-link" title="Snippets" href="/dashboard/snippets">Snippets
</a></li><li class="d-xl-none dropdown header-more">
<a data-toggle="dropdown" href="#">
More
<svg class="caret-down"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#angle-down"></use></svg>
</a>
<div class="dropdown-menu">
<ul>
<li class=""><a title="Activity" href="/dashboard/activity">Activity
</a></li><li class=""><a class="dashboard-shortcuts-milestones" title="Milestones" href="/dashboard/milestones">Milestones
</a></li><li class=""><a class="dashboard-shortcuts-snippets" title="Snippets" href="/dashboard/snippets">Snippets
</a></li>
<li class=""><a title="Instance Statistics" aria-label="Instance Statistics" data-toggle="tooltip" data-placement="bottom" data-container="body" href="/-/instance_statistics">Instance Statistics
</a></li></ul>
</div>
</li>
<li class="hidden">
<a title="Projects" class="dashboard-shortcuts-projects" href="/dashboard/projects">Projects
</a></li>

<li class="d-none d-lg-block d-xl-block"><a title="Instance Statistics" aria-label="Instance Statistics" data-toggle="tooltip" data-placement="bottom" data-container="body" href="/-/instance_statistics"><svg class="s18"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#chart"></use></svg>
</a></li></ul>

</div>
<div class="navbar-collapse collapse">
<ul class="nav navbar-nav">
<li class="header-new dropdown" data-track-event="click_dropdown" data-track-label="new_dropdown">
<a class="header-new-dropdown-toggle has-tooltip qa-new-menu-toggle" title="New..." ref="tooltip" aria-label="New..." data-toggle="dropdown" data-placement="bottom" data-container="body" data-display="static" href="/projects/new"><svg class="s16"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#plus-square"></use></svg>
<svg class="caret-down"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#angle-down"></use></svg>
</a><div class="dropdown-menu dropdown-menu-right">
<ul>
<li class="dropdown-bold-header">
This project
</li>
<li><a href="/developer_web_tools/roku-automated-channel-testing/issues/new">New issue</a></li>
<li><a href="/developer_web_tools/roku-automated-channel-testing/merge_requests/new">New merge request</a></li>
<li><a href="/developer_web_tools/roku-automated-channel-testing/snippets/new">New snippet</a></li>
<li class="divider"></li>
<li class="dropdown-bold-header">GitLab</li>
<li><a class="qa-global-new-project-link" href="/projects/new">New project</a></li>
<li><a class="qa-global-new-snippet-link" href="/snippets/new">New snippet</a></li>
</ul>
</div>
</li>

<li class="nav-item d-none d-sm-none d-md-block m-auto">
<div class="search search-form" data-track-event="activate_form_input" data-track-label="navbar_search">
<form class="form-inline" action="/search" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" /><div class="search-input-container">
<div class="search-input-wrap">
<div class="dropdown" data-url="/search/autocomplete">
<input type="search" name="search" id="search" placeholder="Search or jump to…" class="search-input dropdown-menu-toggle no-outline js-search-dashboard-options" spellcheck="false" tabindex="1" autocomplete="off" data-issues-path="/dashboard/issues" data-mr-path="/dashboard/merge_requests" aria-label="Search or jump to…" />
<button class="hidden js-dropdown-search-toggle" data-toggle="dropdown" type="button"></button>
<div class="dropdown-menu dropdown-select">
<div class="dropdown-content"><ul>
<li class="dropdown-menu-empty-item">
<a>
Loading...
</a>
</li>
</ul>
</div><div class="dropdown-loading"><i aria-hidden="true" data-hidden="true" class="fa fa-spinner fa-spin"></i></div>
</div>
<svg class="s16 search-icon"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#search"></use></svg>
<svg class="s16 clear-icon js-clear-input"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#close"></use></svg>
</div>
</div>
</div>
<input type="hidden" name="group_id" id="group_id" class="js-search-group-options" />
<input type="hidden" name="project_id" id="search_project_id" value="3199" class="js-search-project-options" data-project-path="roku-automated-channel-testing" data-name="roku-automated-channel-testing" data-issues-path="/developer_web_tools/roku-automated-channel-testing/issues" data-mr-path="/developer_web_tools/roku-automated-channel-testing/merge_requests" data-issues-disabled="false" />
<input type="hidden" name="search_code" id="search_code" value="true" />
<input type="hidden" name="repository_ref" id="repository_ref" value="develop" />

<div class="search-autocomplete-opts hide" data-autocomplete-path="/search/autocomplete" data-autocomplete-project-id="3199" data-autocomplete-project-ref="develop"></div>
</form></div>

</li>
<li class="nav-item d-inline-block d-sm-none d-md-none">
<a title="Search" aria-label="Search" data-toggle="tooltip" data-placement="bottom" data-container="body" href="/search?project_id=3199"><svg class="s16"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#search"></use></svg>
</a></li>
<li class="user-counter"><a title="Issues" class="dashboard-shortcuts-issues" aria-label="Issues" data-toggle="tooltip" data-placement="bottom" data-container="body" href="/dashboard/issues?assignee_username=jduval"><svg class="s16"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#issues"></use></svg>
<span class="badge badge-pill green-badge hidden issues-count">
0
</span>
</a></li><li class="user-counter"><a title="Merge requests" class="dashboard-shortcuts-merge_requests" aria-label="Merge requests" data-toggle="tooltip" data-placement="bottom" data-container="body" href="/dashboard/merge_requests?assignee_username=jduval"><svg class="s16"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#git-merge"></use></svg>
<span class="badge badge-pill hidden merge-requests-count">
0
</span>
</a></li><li class="user-counter"><a title="Todos" aria-label="Todos" class="shortcuts-todos" data-toggle="tooltip" data-placement="bottom" data-container="body" href="/dashboard/todos"><svg class="s16"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#todo-done"></use></svg>
<span class="badge badge-pill hidden todos-count">
0
</span>
</a></li><li class="nav-item header-help dropdown">
<a class="header-help-dropdown-toggle" data-toggle="dropdown" href="/help"><svg class="s16"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#question"></use></svg>
<svg class="caret-down"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#angle-down"></use></svg>
</a><div class="dropdown-menu dropdown-menu-right">
<ul>
<li>
<a href="/help">Help</a>
</li>

<li class="divider"></li>
<li>
<a href="https://about.gitlab.com/submit-feedback">Submit feedback</a>
</li>

</ul>

</div>
</li>
<li class="nav-item header-user dropdown" data-track-event="click_dropdown" data-track-label="profile_dropdown">
<a class="header-user-dropdown-toggle" data-toggle="dropdown" href="/jduval"><img width="23" height="23" class="header-user-avatar qa-user-avatar lazy" data-src="https://secure.gravatar.com/avatar/b5109351552b7c8a8722d7d971fc4e8f?s=46&amp;d=identicon" src="data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==" />
<svg class="caret-down"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#angle-down"></use></svg>
</a><div class="dropdown-menu dropdown-menu-right">
<ul>
<li class="current-user">
<div class="user-name bold">
Jonathan Duval
</div>
@jduval
</li>
<li class="divider"></li>
<li>
<div class="js-set-status-modal-trigger" data-has-status="false"></div>
</li>
<li>
<a class="profile-link" data-user="jduval" href="/jduval">Profile</a>
</li>
<li>
<a href="/profile">Settings</a>
</li>
<li class="divider"></li>
<li>
<a class="sign-out-link" href="/users/sign_out">Sign out</a>
</li>
</ul>

</div>
</li>
</ul>
</div>
<button class="navbar-toggler d-block d-sm-none" type="button">
<span class="sr-only">Toggle navigation</span>
<svg class="s12 more-icon js-navbar-toggle-right"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#ellipsis_h"></use></svg>
<svg class="s12 close-icon js-navbar-toggle-left"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#close"></use></svg>
</button>
</div>
</div>
</header>
<div class="js-set-status-modal-wrapper" data-current-emoji="" data-current-message=""></div>

<div class="layout-page page-with-contextual-sidebar">
<div class="nav-sidebar">
<div class="nav-sidebar-inner-scroll">
<div class="context-header">
<a title="roku-automated-channel-testing" href="/developer_web_tools/roku-automated-channel-testing"><div class="avatar-container rect-avatar s40 project-avatar">
<img alt="roku-automated-channel-testing" class="avatar s40 avatar-tile lazy" width="40" height="40" data-src="/uploads/-/system/project/avatar/3199/ic_selenium_driver_512.png" src="data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==" />
</div>
<div class="sidebar-context-title">
roku-automated-channel-testing
</div>
</a></div>
<ul class="sidebar-top-level-items">
<li class="home"><a class="shortcuts-project" href="/developer_web_tools/roku-automated-channel-testing"><div class="nav-icon-container">
<svg><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#home"></use></svg>
</div>
<span class="nav-item-name">
Project
</span>
</a><ul class="sidebar-sub-level-items">
<li class="fly-out-top-item"><a href="/developer_web_tools/roku-automated-channel-testing"><strong class="fly-out-top-item-name">
Project
</strong>
</a></li><li class="divider fly-out-top-item"></li>
<li class=""><a title="Project details" class="shortcuts-project" href="/developer_web_tools/roku-automated-channel-testing"><span>Details</span>
</a></li><li class=""><a title="Activity" class="shortcuts-project-activity qa-activity-link" href="/developer_web_tools/roku-automated-channel-testing/activity"><span>Activity</span>
</a></li><li class=""><a title="Releases" class="shortcuts-project-releases" href="/developer_web_tools/roku-automated-channel-testing/releases"><span>Releases</span>
</a></li>
<li class=""><a title="Cycle Analytics" class="shortcuts-project-cycle-analytics" href="/developer_web_tools/roku-automated-channel-testing/cycle_analytics"><span>Cycle Analytics</span>
</a></li>
</ul>
</li><li class="active"><a class="shortcuts-tree qa-project-menu-repo" href="/developer_web_tools/roku-automated-channel-testing/tree/develop"><div class="nav-icon-container">
<svg><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#doc-text"></use></svg>
</div>
<span class="nav-item-name">
Repository
</span>
</a><ul class="sidebar-sub-level-items">
<li class="fly-out-top-item active"><a href="/developer_web_tools/roku-automated-channel-testing/tree/develop"><strong class="fly-out-top-item-name">
Repository
</strong>
</a></li><li class="divider fly-out-top-item"></li>
<li class="active"><a href="/developer_web_tools/roku-automated-channel-testing/tree/develop">Files
</a></li><li class=""><a href="/developer_web_tools/roku-automated-channel-testing/commits/develop">Commits
</a></li><li class=""><a class="qa-branches-link" href="/developer_web_tools/roku-automated-channel-testing/branches">Branches
</a></li><li class=""><a href="/developer_web_tools/roku-automated-channel-testing/tags">Tags
</a></li><li class=""><a href="/developer_web_tools/roku-automated-channel-testing/graphs/develop">Contributors
</a></li><li class=""><a href="/developer_web_tools/roku-automated-channel-testing/network/develop">Graph
</a></li><li class=""><a href="/developer_web_tools/roku-automated-channel-testing/compare?from=develop&amp;to=develop">Compare
</a></li><li class=""><a href="/developer_web_tools/roku-automated-channel-testing/graphs/develop/charts">Charts
</a></li>
</ul>
</li><li class=""><a class="shortcuts-issues qa-issues-item" href="/developer_web_tools/roku-automated-channel-testing/issues"><div class="nav-icon-container">
<svg><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#issues"></use></svg>
</div>
<span class="nav-item-name">
Issues
</span>
<span class="badge badge-pill count issue_counter">
0
</span>
</a><ul class="sidebar-sub-level-items">
<li class="fly-out-top-item"><a href="/developer_web_tools/roku-automated-channel-testing/issues"><strong class="fly-out-top-item-name">
Issues
</strong>
<span class="badge badge-pill count issue_counter fly-out-badge">
0
</span>
</a></li><li class="divider fly-out-top-item"></li>
<li class=""><a title="Issues" href="/developer_web_tools/roku-automated-channel-testing/issues"><span>
List
</span>
</a></li><li class=""><a title="Boards" href="/developer_web_tools/roku-automated-channel-testing/boards"><span>
Boards
</span>
</a></li><li class=""><a title="Labels" class="qa-labels-link" href="/developer_web_tools/roku-automated-channel-testing/labels"><span>
Labels
</span>
</a></li>
<li class=""><a title="Milestones" class="qa-milestones-link" href="/developer_web_tools/roku-automated-channel-testing/milestones"><span>
Milestones
</span>
</a></li></ul>
</li><li class=""><a class="shortcuts-merge_requests qa-merge-requests-link" href="/developer_web_tools/roku-automated-channel-testing/merge_requests"><div class="nav-icon-container">
<svg><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#git-merge"></use></svg>
</div>
<span class="nav-item-name">
Merge Requests
</span>
<span class="badge badge-pill count merge_counter js-merge-counter">
0
</span>
</a><ul class="sidebar-sub-level-items is-fly-out-only">
<li class="fly-out-top-item"><a href="/developer_web_tools/roku-automated-channel-testing/merge_requests"><strong class="fly-out-top-item-name">
Merge Requests
</strong>
<span class="badge badge-pill count merge_counter js-merge-counter fly-out-badge">
0
</span>
</a></li></ul>
</li><li class=""><a class="shortcuts-pipelines qa-link-pipelines" href="/developer_web_tools/roku-automated-channel-testing/pipelines"><div class="nav-icon-container">
<svg><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#rocket"></use></svg>
</div>
<span class="nav-item-name">
CI / CD
</span>
</a><ul class="sidebar-sub-level-items">
<li class="fly-out-top-item"><a href="/developer_web_tools/roku-automated-channel-testing/pipelines"><strong class="fly-out-top-item-name">
CI / CD
</strong>
</a></li><li class="divider fly-out-top-item"></li>
<li class=""><a title="Pipelines" class="shortcuts-pipelines" href="/developer_web_tools/roku-automated-channel-testing/pipelines"><span>
Pipelines
</span>
</a></li><li class=""><a title="Jobs" class="shortcuts-builds" href="/developer_web_tools/roku-automated-channel-testing/-/jobs"><span>
Jobs
</span>
</a></li><li class=""><a title="Schedules" class="shortcuts-builds" href="/developer_web_tools/roku-automated-channel-testing/pipeline_schedules"><span>
Schedules
</span>
</a></li><li class=""><a title="Charts" class="shortcuts-pipelines-charts" href="/developer_web_tools/roku-automated-channel-testing/pipelines/charts"><span>
Charts
</span>
</a></li></ul>
</li><li class=""><a class="shortcuts-operations qa-link-operations" href="/developer_web_tools/roku-automated-channel-testing/environments/metrics"><div class="nav-icon-container">
<svg><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#cloud-gear"></use></svg>
</div>
<span class="nav-item-name">
Operations
</span>
</a><ul class="sidebar-sub-level-items">
<li class="fly-out-top-item"><a href="/developer_web_tools/roku-automated-channel-testing/environments/metrics"><strong class="fly-out-top-item-name">
Operations
</strong>
</a></li><li class="divider fly-out-top-item"></li>
<li class=""><a title="Metrics" class="shortcuts-metrics" href="/developer_web_tools/roku-automated-channel-testing/environments/metrics"><span>
Metrics
</span>
</a></li>
<li class=""><a title="Environments" class="shortcuts-environments qa-operations-environments-link" href="/developer_web_tools/roku-automated-channel-testing/environments"><span>
Environments
</span>
</a></li><li class=""><a title="Error Tracking" class="shortcuts-tracking qa-operations-tracking-link" href="/developer_web_tools/roku-automated-channel-testing/error_tracking"><span>
Error Tracking
</span>
</a></li></ul>
</li><li class=""><a class="shortcuts-wiki qa-wiki-link" href="/developer_web_tools/roku-automated-channel-testing/wikis/home"><div class="nav-icon-container">
<svg><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#book"></use></svg>
</div>
<span class="nav-item-name">
Wiki
</span>
</a><ul class="sidebar-sub-level-items is-fly-out-only">
<li class="fly-out-top-item"><a href="/developer_web_tools/roku-automated-channel-testing/wikis/home"><strong class="fly-out-top-item-name">
Wiki
</strong>
</a></li></ul>
</li><li class=""><a class="shortcuts-snippets" href="/developer_web_tools/roku-automated-channel-testing/snippets"><div class="nav-icon-container">
<svg><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#snippet"></use></svg>
</div>
<span class="nav-item-name">
Snippets
</span>
</a><ul class="sidebar-sub-level-items is-fly-out-only">
<li class="fly-out-top-item"><a href="/developer_web_tools/roku-automated-channel-testing/snippets"><strong class="fly-out-top-item-name">
Snippets
</strong>
</a></li></ul>
</li><li class=""><a title="Members" class="shortcuts-tree" href="/developer_web_tools/roku-automated-channel-testing/settings/members"><div class="nav-icon-container">
<svg><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#users"></use></svg>
</div>
<span class="nav-item-name">
Members
</span>
</a><ul class="sidebar-sub-level-items is-fly-out-only">
<li class="fly-out-top-item"><a href="/developer_web_tools/roku-automated-channel-testing/project_members"><strong class="fly-out-top-item-name">
Members
</strong>
</a></li></ul>
</li><a class="toggle-sidebar-button js-toggle-sidebar" role="button" title="Toggle sidebar" type="button">
<svg class="icon-angle-double-left"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#angle-double-left"></use></svg>
<svg class="icon-angle-double-right"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#angle-double-right"></use></svg>
<span class="collapse-text">Collapse sidebar</span>
</a>
<button name="button" type="button" class="close-nav-button"><svg class="s16"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#close"></use></svg>
<span class="collapse-text">Close sidebar</span>
</button>
<li class="hidden">
<a title="Activity" class="shortcuts-project-activity" href="/developer_web_tools/roku-automated-channel-testing/activity"><span>
Activity
</span>
</a></li>
<li class="hidden">
<a title="Network" class="shortcuts-network" href="/developer_web_tools/roku-automated-channel-testing/network/develop">Graph
</a></li>
<li class="hidden">
<a title="Charts" class="shortcuts-repository-charts" href="/developer_web_tools/roku-automated-channel-testing/graphs/develop/charts">Charts
</a></li>
<li class="hidden">
<a class="shortcuts-new-issue" href="/developer_web_tools/roku-automated-channel-testing/issues/new">Create a new issue
</a></li>
<li class="hidden">
<a title="Jobs" class="shortcuts-builds" href="/developer_web_tools/roku-automated-channel-testing/-/jobs">Jobs
</a></li>
<li class="hidden">
<a title="Commits" class="shortcuts-commits" href="/developer_web_tools/roku-automated-channel-testing/commits/develop">Commits
</a></li>
<li class="hidden">
<a title="Issue Boards" class="shortcuts-issue-boards" href="/developer_web_tools/roku-automated-channel-testing/boards">Issue Boards</a>
</li>
</ul>
</div>
</div>

<div class="content-wrapper">

<div class="mobile-overlay"></div>
<div class="alert-wrapper">






<nav class="breadcrumbs container-fluid container-limited" role="navigation">
<div class="breadcrumbs-container">
<button name="button" type="button" class="toggle-mobile-nav"><span class="sr-only">Open sidebar</span>
<i aria-hidden="true" data-hidden="true" class="fa fa-bars"></i>
</button><div class="breadcrumbs-links js-title-container">
<ul class="list-unstyled breadcrumbs-list js-breadcrumbs-list">
<li><a class="group-path breadcrumb-item-text js-breadcrumb-item-text " href="/developer_web_tools">developer_web_tools</a><svg class="s8 breadcrumbs-list-angle"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#angle-right"></use></svg></li> <li><a href="/developer_web_tools/roku-automated-channel-testing"><img alt="roku-automated-channel-testing" class="avatar-tile lazy" width="15" height="15" data-src="/uploads/-/system/project/avatar/3199/ic_selenium_driver_512.png" src="data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==" /><span class="breadcrumb-item-text js-breadcrumb-item-text">roku-automated-channel-testing</span></a><svg class="s8 breadcrumbs-list-angle"><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#angle-right"></use></svg></li>

<li>
<h2 class="breadcrumbs-sub-title"><a href="/developer_web_tools/roku-automated-channel-testing/blob/develop/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot">Repository</a></h2>
</li>
</ul>
</div>

</div>
</nav>

<div class="flash-container flash-container-page">
</div>

<div class="d-flex"></div>
</div>
<div class=" ">
<div class="content" id="content-body">
<div class="js-signature-container" data-signatures-path="/developer_web_tools/roku-automated-channel-testing/commits/b1dde470963eaac85647abcdb5ef5cc14b7b87c9/signatures"></div>
<div class="container-fluid container-limited">

<div class="tree-holder" id="tree-holder">
<div class="nav-block">
<div class="tree-ref-container">
<div class="tree-ref-holder">
<form class="project-refs-form" action="/developer_web_tools/roku-automated-channel-testing/refs/switch" accept-charset="UTF-8" method="get"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="destination" id="destination" value="blob" />
<input type="hidden" name="path" id="path" value="RobotLibrary/Tests/test_7_EpisodePicker_utils.robot" />
<div class="dropdown">
<button class="dropdown-menu-toggle js-project-refs-dropdown qa-branches-select" type="button" data-toggle="dropdown" data-selected="develop" data-ref="develop" data-refs-url="/developer_web_tools/roku-automated-channel-testing/refs?sort=updated_desc" data-field-name="ref" data-submit-form-on-click="true" data-visit="true"><span class="dropdown-toggle-text ">develop</span><i aria-hidden="true" data-hidden="true" class="fa fa-chevron-down"></i></button>
<div class="dropdown-menu dropdown-menu-paging dropdown-menu-selectable git-revision-dropdown qa-branches-dropdown">
<div class="dropdown-page-one">
<div class="dropdown-title"><span>Switch branch/tag</span><button class="dropdown-title-button dropdown-menu-close" aria-label="Close" type="button"><i aria-hidden="true" data-hidden="true" class="fa fa-times dropdown-menu-close-icon"></i></button></div>
<div class="dropdown-input"><input type="search" id="" class="dropdown-input-field" placeholder="Search branches and tags" autocomplete="off" /><i aria-hidden="true" data-hidden="true" class="fa fa-search dropdown-input-search"></i><i aria-hidden="true" data-hidden="true" role="button" class="fa fa-times dropdown-input-clear js-dropdown-input-clear"></i></div>
<div class="dropdown-content"></div>
<div class="dropdown-loading"><i aria-hidden="true" data-hidden="true" class="fa fa-spinner fa-spin"></i></div>
</div>
</div>
</div>
</form>
</div>
<ul class="breadcrumb repo-breadcrumb">
<li class="breadcrumb-item">
<a href="/developer_web_tools/roku-automated-channel-testing/tree/develop">roku-automated-channel-testing
</a></li>
<li class="breadcrumb-item">
<a href="/developer_web_tools/roku-automated-channel-testing/tree/develop/RobotLibrary">RobotLibrary</a>
</li>
<li class="breadcrumb-item">
<a href="/developer_web_tools/roku-automated-channel-testing/tree/develop/RobotLibrary/Tests">Tests</a>
</li>
<li class="breadcrumb-item">
<a href="/developer_web_tools/roku-automated-channel-testing/blob/develop/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot"><strong>test_7_EpisodePicker_utils.robot</strong>
</a></li>
</ul>
</div>
<div class="tree-controls">
<a class="btn shortcuts-find-file" rel="nofollow" href="/developer_web_tools/roku-automated-channel-testing/find_file/develop"><i aria-hidden="true" data-hidden="true" class="fa fa-search"></i>
<span>Find file</span>
</a>
<div class="btn-group" role="group"><a class="btn js-blob-blame-link" href="/developer_web_tools/roku-automated-channel-testing/blame/develop/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot">Blame</a><a class="btn" href="/developer_web_tools/roku-automated-channel-testing/commits/develop/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot">History</a><a class="btn js-data-file-blob-permalink-url" href="/developer_web_tools/roku-automated-channel-testing/blob/558c402c090d08dc20b50176f24a102a5eb86e99/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot">Permalink</a></div>
</div>
</div>

<div class="info-well d-none d-sm-block">
<div class="well-segment">
<ul class="blob-commit-info">
<li class="commit flex-row js-toggle-container" id="commit-b1dde470">
<div class="avatar-cell d-none d-sm-block">
<a href="mailto:andrii.kovtko@synapse.com"><img alt="Andrii Kovtko&#39;s avatar" src="https://secure.gravatar.com/avatar/193066fbf5835d9218c2f37216be9215?s=72&amp;d=identicon" class="avatar s36 d-none d-sm-inline" title="Andrii Kovtko" /></a>
</div>
<div class="commit-detail flex-list">
<div class="commit-content qa-commit-content">
<a class="commit-row-message item-title" href="/developer_web_tools/roku-automated-channel-testing/commit/b1dde470963eaac85647abcdb5ef5cc14b7b87c9">added get child nodes keyword</a>
<span class="commit-row-message d-inline d-sm-none">
&middot;
b1dde470
</span>
<div class="committer">
<a class="commit-author-link" href="mailto:andrii.kovtko@synapse.com">Andrii Kovtko</a> authored <time class="js-timeago" title="Jul 10, 2020 6:09am" datetime="2020-07-10T13:09:43Z" data-toggle="tooltip" data-placement="bottom" data-container="body">Jul 10, 2020</time>
</div>
</div>
<div class="commit-actions flex-row">

<div class="js-commit-pipeline-status" data-endpoint="/developer_web_tools/roku-automated-channel-testing/commit/b1dde470963eaac85647abcdb5ef5cc14b7b87c9/pipelines?ref=develop"></div>
<div class="commit-sha-group d-none d-sm-flex">
<div class="label label-monospace monospace">
b1dde470
</div>
<button class="btn btn btn-default" data-toggle="tooltip" data-placement="bottom" data-container="body" data-title="Copy commit SHA to clipboard" data-class="btn btn-default" data-clipboard-text="b1dde470963eaac85647abcdb5ef5cc14b7b87c9" type="button" title="Copy commit SHA to clipboard" aria-label="Copy commit SHA to clipboard"><svg><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#duplicate"></use></svg></button>

</div>
</div>
</div>
</li>

</ul>
</div>


</div>
<div class="blob-content-holder" id="blob-content-holder">
<article class="file-holder">
<div class="js-file-title file-title-flex-parent">
<div class="file-header-content">
<i aria-hidden="true" data-hidden="true" class="fa fa-file-text-o fa-fw"></i>
<strong class="file-title-name qa-file-title-name">
test_7_EpisodePicker_utils.robot
</strong>
<button class="btn btn-clipboard btn-transparent prepend-left-5" data-toggle="tooltip" data-placement="bottom" data-container="body" data-class="btn-clipboard btn-transparent prepend-left-5" data-title="Copy file path to clipboard" data-clipboard-text="{&quot;text&quot;:&quot;RobotLibrary/Tests/test_7_EpisodePicker_utils.robot&quot;,&quot;gfm&quot;:&quot;`RobotLibrary/Tests/test_7_EpisodePicker_utils.robot`&quot;}" type="button" title="Copy file path to clipboard" aria-label="Copy file path to clipboard"><svg><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#duplicate"></use></svg></button>
<small>
2.39 KB
</small>
</div>

<div class="file-actions">

<div class="btn-group" role="group"><button class="btn btn btn-sm js-copy-blob-source-btn" data-toggle="tooltip" data-placement="bottom" data-container="body" data-class="btn btn-sm js-copy-blob-source-btn" data-title="Copy source to clipboard" data-clipboard-target=".blob-content[data-blob-id=&#39;420412138c90e33d6fc7c79c81c5aaa1ed33c2c4&#39;]" type="button" title="Copy source to clipboard" aria-label="Copy source to clipboard"><svg><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#duplicate"></use></svg></button><a class="btn btn-sm has-tooltip" target="_blank" rel="noopener noreferrer" title="Open raw" data-container="body" href="/developer_web_tools/roku-automated-channel-testing/raw/develop/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot"><i aria-hidden="true" data-hidden="true" class="fa fa-file-code-o"></i></a><a download="RobotLibrary/Tests/test_7_EpisodePicker_utils.robot" class="btn btn-sm has-tooltip" target="_blank" rel="noopener noreferrer" title="Download" data-container="body" href="/developer_web_tools/roku-automated-channel-testing/raw/develop/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot?inline=false"><svg><use xlink:href="/assets/icons-09fdf2c02921bad2ec7257465016a755f359ab7b598e5fe42c22381fe1a25045.svg#download"></use></svg></a></div>
<div class="btn-group" role="group">
<a class="btn js-edit-blob  btn-sm" href="/developer_web_tools/roku-automated-channel-testing/edit/develop/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot">Edit</a><a class="btn btn-default btn-sm" href="/-/ide/project/developer_web_tools/roku-automated-channel-testing/edit/develop/-/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot">Web IDE</a><button name="button" type="submit" class="btn btn-default" data-target="#modal-upload-blob" data-toggle="modal">Replace</button><button name="button" type="submit" class="btn btn-remove" data-target="#modal-remove-blob" data-toggle="modal">Delete</button></div>
</div>
</div>
<div class="js-file-fork-suggestion-section file-fork-suggestion hidden">
<span class="file-fork-suggestion-note">
You're not allowed to
<span class="js-file-fork-suggestion-section-action">
edit
</span>
files in this project directly. Please fork this project,
make your changes there, and submit a merge request.
</span>
<a class="js-fork-suggestion-button btn btn-grouped btn-inverted btn-success" rel="nofollow" data-method="post" href="/developer_web_tools/roku-automated-channel-testing/blob/develop/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot">Fork</a>
<button class="js-cancel-fork-suggestion-button btn btn-grouped" type="button">
Cancel
</button>
</div>



<div class="blob-viewer" data-type="simple" data-url="/developer_web_tools/roku-automated-channel-testing/blob/develop/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot?format=json&amp;viewer=simple">
<div class="text-center prepend-top-default append-bottom-default">
<i aria-hidden="true" aria-label="Loading content…" class="fa fa-spinner fa-spin fa-2x qa-spinner"></i>
</div>

</div>


</article>
</div>

<div class="modal" id="modal-remove-blob">
<div class="modal-dialog">
<div class="modal-content">
<div class="modal-header">
<h3 class="page-title">Delete test_7_EpisodePicker_utils.robot</h3>
<button aria-label="Close" class="close" data-dismiss="modal" type="button">
<span aria-hidden="true">&times;</span>
</button>
</div>
<div class="modal-body">
<form class="js-delete-blob-form js-quick-submit js-requires-input" action="/developer_web_tools/roku-automated-channel-testing/blob/develop/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="_method" value="delete" /><input type="hidden" name="authenticity_token" value="el5xl4N9swPOkxIoQPh1NL0ZHcbZ38TxmcCIYxGRuyn0Xplo1sGXPJ4ifLVIgFSv+xIH07RWOlt74Lv7bKQPkg==" /><div class="form-group row commit_message-group">
<label class="col-form-label col-sm-2" for="commit_message-fd4297b15f855baaeb9b41aa485f8ef9">Commit message
</label><div class="col-sm-10">
<div class="commit-message-container">
<div class="max-width-marker"></div>
<textarea name="commit_message" id="commit_message-fd4297b15f855baaeb9b41aa485f8ef9" class="form-control js-commit-message" placeholder="Delete test_7_EpisodePicker_utils.robot" required="required" rows="3">
Delete test_7_EpisodePicker_utils.robot</textarea>
</div>
</div>
</div>

<div class="form-group row branch">
<label class="col-form-label col-sm-2" for="branch_name">Target Branch</label>
<div class="col-sm-10">
<input type="text" name="branch_name" id="branch_name" value="develop" required="required" class="form-control js-branch-name ref-name" />
<div class="js-create-merge-request-container">
<div class="form-check prepend-top-8">
<input type="checkbox" name="create_merge_request" id="create_merge_request-abe095d1b9dba4ee9699b74df5bc4b6e" value="1" class="js-create-merge-request form-check-input" checked="checked" />
<label class="form-check-label" for="create_merge_request-abe095d1b9dba4ee9699b74df5bc4b6e">Start a <strong>new merge request</strong> with these changes
</label></div>

</div>
</div>
</div>
<input type="hidden" name="original_branch" id="original_branch" value="develop" class="js-original-branch" />

<div class="form-group row">
<div class="offset-sm-2 col-sm-10">
<button name="button" type="submit" class="btn btn-remove btn-remove-file">Delete file</button>
<a class="btn btn-cancel" data-dismiss="modal" href="#">Cancel</a>
</div>
</div>
</form></div>
</div>
</div>
</div>

<div class="modal" id="modal-upload-blob">
<div class="modal-dialog modal-lg">
<div class="modal-content">
<div class="modal-header">
<h3 class="page-title">Replace test_7_EpisodePicker_utils.robot</h3>
<button aria-label="Close" class="close" data-dismiss="modal" type="button">
<span aria-hidden="true">&times;</span>
</button>
</div>
<div class="modal-body">
<form class="js-quick-submit js-upload-blob-form" data-method="put" action="/developer_web_tools/roku-automated-channel-testing/update/develop/RobotLibrary/Tests/test_7_EpisodePicker_utils.robot" accept-charset="UTF-8" method="post"><input name="utf8" type="hidden" value="&#x2713;" /><input type="hidden" name="_method" value="put" /><input type="hidden" name="authenticity_token" value="CxFBr4Eh8zN0Dv80EP+pqawfCLRpKcGXchpOp5980DCFEalQ1J3XDCS/kakYh4gy6hQSoQSgPz2QOn0/4klkiw==" /><div class="dropzone">
<div class="dropzone-previews blob-upload-dropzone-previews">
<p class="dz-message light">
Attach a file by drag &amp; drop or <a class="markdown-selector" href="#">click to upload</a>
</p>
</div>
</div>
<br>
<div class="dropzone-alerts alert alert-danger data" style="display:none"></div>
<div class="form-group row commit_message-group">
<label class="col-form-label col-sm-2" for="commit_message-e495a1ecb7c543343d7a58ae01691efb">Commit message
</label><div class="col-sm-10">
<div class="commit-message-container">
<div class="max-width-marker"></div>
<textarea name="commit_message" id="commit_message-e495a1ecb7c543343d7a58ae01691efb" class="form-control js-commit-message" placeholder="Replace test_7_EpisodePicker_utils.robot" required="required" rows="3">
Replace test_7_EpisodePicker_utils.robot</textarea>
</div>
</div>
</div>

<div class="form-group row branch">
<label class="col-form-label col-sm-2" for="branch_name">Target Branch</label>
<div class="col-sm-10">
<input type="text" name="branch_name" id="branch_name" value="develop" required="required" class="form-control js-branch-name ref-name" />
<div class="js-create-merge-request-container">
<div class="form-check prepend-top-8">
<input type="checkbox" name="create_merge_request" id="create_merge_request-bf2ff771cb0044d893183f548229d37a" value="1" class="js-create-merge-request form-check-input" checked="checked" />
<label class="form-check-label" for="create_merge_request-bf2ff771cb0044d893183f548229d37a">Start a <strong>new merge request</strong> with these changes
</label></div>

</div>
</div>
</div>
<input type="hidden" name="original_branch" id="original_branch" value="develop" class="js-original-branch" />

<div class="form-actions">
<button name="button" type="button" class="btn btn-success btn-upload-file" id="submit-all"><i aria-hidden="true" data-hidden="true" class="fa fa-spin fa-spinner js-loading-icon hidden"></i>
Replace file
</button><a class="btn btn-cancel" data-dismiss="modal" href="#">Cancel</a>

</div>
</form></div>
</div>
</div>
</div>

</div>
</div>

</div>
</div>
</div>
</div>



</body>
</html>

