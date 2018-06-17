var app = angular.module("app",[]);
app.controller("ctrl", function($scope,$http) {

    $scope.Error=0
    $scope.SaveTime=function(){
        if(typeof $scope.From=="undefined" || typeof $scope.To=="undefined" || 
            typeof $scope.Statement_encoded =="undefined"){
            $scope.ShowNewStatementMenu=0;
            $scope.Error="You Can't Leave Values Empty!";
            return;
        }

        if ($scope.To.split(':')[0] > 23 || $scope.From.split(':')[0] > 23){
            $scope.Error="You Can't have Hours after 23 (hint: try 0 instead of 24)"
            return;
        }
        if ($scope.To.split(':')[1] > 60 || $scope.From.split(':')[1] > 60){
            $scope.Error="You Can't have Minutes after 60"
            return;
        }
        datefrom=new Date();
        datefrom.setHours($scope.From.split(':')[0]);
        datefrom.setMinutes($scope.From.split(':')[1]);

        dateto=new Date();
        dateto.setHours($scope.To.split(':')[0]);
        dateto.setMinutes($scope.To.split(':')[1]);

        dateFromUnix=datefrom.getTime()/1000
        dateToUnix=dateto.getTime()/1000
        $http.get("/deposit/"+dateFromUnix+"/"+dateToUnix+"/"+$scope.Statement_encoded).then(
        function(response) {
          if (response.data=="ERROR"){
                $scope.Error="Unable to Save Time Check the Service Logs";
        }});
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