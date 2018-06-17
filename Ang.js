var app = angular.module("app",[]);
app.controller("ctrl", function($scope) {
    $scope.SaveTime=function(){
        console.log($scope.From)
        console.log($scope.To)
        console.log($scope.Statement_encoded)
    }
    $scope.SaveStatement=function(){
        console.log($scope.Statement_encoded_input)
        console.log($scope.Statement_real)
    }
});