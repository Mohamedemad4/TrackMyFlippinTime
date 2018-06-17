var app = angular.module("app",[]);
app.controller("ctrl", function($scope) {
    $scope.h1 ='Joke`s on you Hater';
    $scope.fg=0;
    $scope.SaveTime=function(){
        console.log($scope.From)
    }
    $scope.SaveStatement=function(){
        console.log($scope.Statement_encoded_input)
    }
});