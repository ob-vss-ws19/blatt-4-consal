echo "Start User"
docker run terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-consal:PR-3-userservice &
sleep 0.2
echo "Start Movieservice"
docker run terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-consal:PR-3-movieservice &
sleep 0.2
echo "Start Cinemahall"
docker run terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-consal:PR-3-cinemahallservice &
sleep 0.2
echo "Start Show"
docker run terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-consal:PR-3-showservice &
sleep 0.2
echo "Start Reservation"
docker run terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-consal:PR-3-reservationservice &
