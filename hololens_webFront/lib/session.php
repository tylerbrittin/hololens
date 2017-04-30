<?php

session_start();

const SESSION_KEY = 'user';

/**
 * Check if an user is signed in
 * @return bool
 */
function is_logged_in() {
    return array_key_exists(SESSION_KEY, $_SESSION) && array_key_exists('username', $_SESSION[SESSION_KEY]);
}

/**
 * Get the details of logged in user
 * @return array|null
 */
function get_logged_in_user() {
    return is_logged_in() ? $_SESSION[SESSION_KEY] : null;
}

/**
 * Log in as a specific user
 * @param string $username
 * @param string $email
 */
function log_in($username, $email) {
    $_SESSION[SESSION_KEY] = [
        'username' => $username,
        'email' => $email
    ];
}

/**
 * Log out of the existing session
 */
function log_out() {
    unset($_SESSION[SESSION_KEY]);
}