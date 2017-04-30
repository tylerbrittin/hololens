<?php

/**
 * Redirect the user to the provided URL
 * @param string $url
 */
function redirect($url) {
    header('Location: '.$url);
    exit();
}

/**
 * Check if the current request is a POST request
 * @return bool
 */
function is_post_request() {
    return $_SERVER['REQUEST_METHOD'] == 'POST';
}