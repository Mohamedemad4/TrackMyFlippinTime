var app = angular.module("app",[]);
app.controller("ctrl", function($scope,$http) {
    $scope.Error=0
    $scope.SaveTime=function(){
        console.log($scope.From)
        console.log($scope.To)
        console.log($scope.Statement_encoded)
    }

    $scope.SaveStatement=function(){
        if(typeof $scope.Statemnet_real=="undefined" || typeof $scope.Statement_encoded_input=="undefined"){
            $scope.ShowNewStatementMenu=0;
            $scope.Error="You Can't Leave Values Empty!";
            return;
        }
        $http.get("/newstatement/"+$scope.Statemnet_real+"/"+$scope.Statement_encoded_input).then(
        function(response) {
          if (response.data=="ERROR"){
                $scope.Error="Unable to Save New Statement Check the Service Logs";
        }});
    }

});