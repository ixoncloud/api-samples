
import * as appSettings from "tns-core-modules/application-settings";


// import { ActivityService } from "./_services/activity.service";

@JavaProxy("com.tns.MyWidget")
export class WidgetComponent extends android.appwidget.AppWidgetProvider {

    // public activityService: ActivityService

    public test:string

    onUpdate(context, appWidgetManager, appWidgetIds): void {
      var appWidgetsLen = appWidgetIds.length

      for (let i = 0; i < appWidgetsLen; i++) {
        this.updateWidget(context, appWidgetManager, appWidgetIds, appWidgetIds[i]);
        
      }
    }
    
    updateWidget(context, appWidgetManager, appWidgetIds, widgetId){

      console.log("update called");
      
      this.test = appSettings.getString("test")
      console.log(this.test);
      

      let layout:any = context.getResources().getIdentifier("appwidget", "layout", context.getPackageName())
      let resourceId:any = context.getResources().getIdentifier("text2", "id", context.getPackageName())
      let imageId:any = context.getResources().getIdentifier("image", "id", context.getPackageName())
      let image:any = context.getResources().getIdentifier("ns", "drawable", context.getPackageName())

      var views: android.widget.RemoteViews = new android.widget.RemoteViews(context.getPackageName(), layout);
      
      views.setTextViewText(resourceId, this.test);

      // views.setImageViewResource(imageId, image)
      android.appwidget.AppWidgetManager.getInstance(context).updateAppWidget(
        new android.content.ComponentName(context, WidgetComponent.class), views
      );
    }

    createImg(): any{

    }
}
