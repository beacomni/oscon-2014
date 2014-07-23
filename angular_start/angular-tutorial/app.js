angular.module('musicPlayer',[]).
  controller('PlaylistCtrl', function($scope,$filter,$http) {

   	$scope.currentSong = "bohemain rhapsody";
 	$scope.playlist = [];

 	$scope.playlistSearch = "";

 	$scope.currentListName = "";
 	$scope.savedPlaylists = [];
 	$scope.currentPlaylist;

    $http.get('/playlists.json').
 	   success(function(data) {
		   data.forEach(function(pList) {
			   $scope.savedPlaylists.push(pList);
		   })
 		   $scope.savedPlaylists.push(data);
 	   }).
        error(function() {
 		   alert('error');
 	   });		   

 	$scope.addSong = function() {

 		var lowercaseFilter = $filter('lowercase');

 		$scope.playlist.push(lowercaseFilter($scope.currentSong));
 		$scope.currentSong = "";
 	}

 	$scope.removeSong = function(song) {
 		var idx = $scope.playlist.indexOf(song);
 		$scope.playlist.splice(idx,1);
 	}

 	$scope.savePlaylist = function() {
 		var orderBy = $filter('orderBy');

 		$scope.savedPlaylists.push({
          name: $scope.currentListName, 
 		  playlist : $scope.playlist});
 		$scope.savedPlaylists = orderBy($scope.savedPlaylists,'name');
 		$scope.currentListName = '';
 		$scope.currentSong = '';
 		$scope.playlist = [];
 		//$scope.name = "";
 	}

 	$scope.setCurrentPlaylist = function(playlist) {
		if (playlist.file !== undefined){
			$http.get('/' + playlist.file).
				success(function(data) {
					$scope.currentPlaylist = data;
				})
		}
 		$scope.currentPlaylist = playlist;
 	}
   }	
)


