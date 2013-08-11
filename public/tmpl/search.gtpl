<html>
   <head>
      <title>Search</title>
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <link href="static/bootstrap/css/bootstrap.css" rel="stylesheet" media="screen">
      <link href="static/bootstrap/css/bootstrap-responsive.css" rel="stylesheet" media="screen">
   </head>
   <body>
      <div class="row-fluid">
         <div class="span5 offset1">
            <div class="alert alert-error" id="alert" style="display:none">
            </div>
            <form id="distanceForm" class="form-horizontal">
               <fieldset>
                  <legend>Distance</legend>
                  <div class="row">
                     <div class="span7">
                        <div class="control-group">
                           <label class="control-label" for="from">From: </label>
                           <div class="controls">
                              <input type="text" id="from" value="ta22 9rt" placeholder="From" autofocus="true" name="from" class="span12"/>
                           </div>
                        </div>
                     </div>
                     <div class="span5" id="address1" style="display:none">
                     </div>
                  </div>
                  <div class="row">
                     <div class="span7">
                        <div class="control-group">
                           <label class="control-label" for="to">To: </label>
                           <div class="controls">
                              <input type="text" id="to" value="ex16 4qa" placeholder="To"  class="span12"/>
                           </div>
                        </div>
                     </div>
                     <div class="span5" id="address2" style="display:none">
                     </div>
                  </div>
                  <div class="row">
                  	<div class="span7"></div>
                  	<div class="span5">
								<div class="row" id="distanceResult"  style="display:none">
				               <label class="control-label span4">Road distance:</label>
				               <div class="label label-info">3367 meters</div>
				            </div>
				            <div class="row" id="distanceResult2"  style="display:none">
									<label class="control-label span4">Line distance:</label>
									<div class="label label-info">3367 meters</div>
								</div>
                  	</div>
                  </div>
                  <div class="form-actions">
                     <input type="submit" value="Go" class="btn btn-primary pull-right">
                  </div>
               </fieldset>
            </form>
         </div>
      </div>
      <ul id="info">
      </ul>
      <script src="static/js/jquery.js"></script>
      <script src="static/bootstrap/js/bootstrap.js"></script>
      <script src="static/js/main.js"></script>
   </body>
</html>
