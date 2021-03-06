<?php

require_once __DIR__.'/../libs/httpful.phar';
const API = "http://40.71.214.175:5073/";

function check_username_availability($username) {
    $uri = API."checkusername/".urlencode($username);
    $response = \Httpful\Request::get($uri)->send();
    return !$response->body->Taken;
}

function get_user_info($username) {
    $uri = API."getuserinfo/".urlencode($username);
    $response = \Httpful\Request::get($uri)->send();
    return $response->body;
}

function verify_user($username, $password, $get_info = true) {
    $password = md5($password);
    $info = get_user_info($username);
    if(!strlen($info->Username)) return false; //user does not exist

    if($get_info) {
        return $password === $info->Password ? $info : null;
    } else {
        return $password === $info->Password;
    }
}

function create_user($user) {
    $uri = API."adduser";
    $response = \Httpful\Request::post($uri)
        ->sendsJson()
        ->body(json_encode([
            'Username' => $user['username'],
            'Password' => md5($user['password']),
            'Firstname' => $user['first_name'],
            'Lastname' => $user['last_name'],
            'Email' => $user['email']
        ]))
        ->send();
    return !$response->hasErrors();
}

function delete_item($seller, $category, $id) {
    $uri = API."deleteitem";
    $response = \Httpful\Request::post($uri)
        ->sendsJson()
        ->body(json_encode([
                'Seller' => $seller,
                'Category' => $category,
                'id' => $id
            ]))
        ->send();
        return !$response->hasErrors();
}

function edit_item($email, $seller, $category, $item, $itemdesc, $price, $id, $model, $texture) {
    $uri = API."edititem";
    $response = \Httpful\Request::post($uri)
        ->sendsJson()
        ->body(json_encode([
                'Id' => $id,
                'Category' => $category,
                'Seller' => $seller,
                'Price' => $price,
                'Model' => $model,
                'Texture' => $texture,
                'Email' => $email,
                'Item' => $item,
                'Itemdesc' => $itemdesc,

            ]))
        ->send();
    return !$response->hasErrors();
}