import comp100.*;
import java.util.*;
import java.io.*;
import javax.swing.*;
import java.awt.event.*;
import java.awt.BorderLayout;
import java.awt.Color;

public class Mine implements MouseListener, WindowListener
{
    private int width = 16; private int height = 16;
    private int[][] array = new int[width][height];
    private int gridw = (800/width) * width;
    private int gridh = (800/height) * height;
    private boolean[][] revealed = new boolean[width][height];
    private boolean[][] flagged = new boolean[width][height];
    private JFrame frame = new JFrame("Mine");
    private DrawingCanvas canvas = new DrawingCanvas();
    private int mines = 40;
    private boolean clicked;    //checks for first click
    private boolean win = true;
    private File records;
    private int wins = 0;
    private int losses = 0;
    private double awt = 0;
    private double bwt = 100;
    private long currentStartTime = 0;
    
    public void setup()
    {
        frame.setSize(gridw + 50,gridh + 50);
        canvas.addMouseListener(this);
        frame.getContentPane().add(canvas, BorderLayout.CENTER);
        frame.setDefaultCloseOperation(JFrame.DO_NOTHING_ON_CLOSE);
        frame.setVisible(true);
        frame.addWindowListener(this);
        
        records = new File("records.txt");
        try
        {
        	Scanner sc = new Scanner(records);
        	//File has formal wins, losses, average win time, best win time
        	wins = sc.nextInt();
        	losses = sc.nextInt();
        	awt = sc.nextDouble();
        	bwt = sc.nextDouble();
        	sc.close();
        }
        catch (IOException e)
        {
        	//file doesn't exist
        };
        
        start();
    }
    
    public void start()
    {
        canvas.clear();
        revealed = new boolean[width][height];
        flagged = new boolean[width][height];
        array = new int[width][height];
        clicked = false;

        //draws the board
        canvas.setColor(Color.black);
        canvas.fillRect(0, 0, gridw, gridh);
        if (win)
        {
            canvas.setColor(Color.blue);
        }
        else
        {
            canvas.setColor(Color.red);
        }
        
        for (int x = 0; x < width; x++)
        {
            for (int y = 0; y < height; y++)
            {
                canvas.fillRect(x*gridw/width, y*gridh/height, gridw/width - 2, gridh/height - 2);
            }
        }
    }
    
    private void reveal(int blockX, int blockY)
    {
        if (!clicked)
        {
        	//start the timer.
        	currentStartTime = System.currentTimeMillis();
            clicked = true;
            createMine(blockX, blockY);
        }
            
        if (blockX >= 0 && blockX < width && blockY >= 0 && blockY < height && !revealed[blockX][blockY] && !flagged[blockX][blockY])
        {
            revealed[blockX][blockY] = true;
            
            if (array[blockX][blockY] > 0)
            {
                canvas.setColor(Color.white);
                canvas.fillRect(blockX*gridw/width, blockY*gridh/height, gridw/width - 2, gridh/height - 2);
                canvas.setColor(Color.black);
                canvas.drawString( array[blockX][blockY] + "", (blockX + 0.5)*gridw/width, (blockY + 0.5)*gridh/height);
            }
            else if (array[blockX][blockY] == -1)
            {
                canvas.setColor(Color.red);
                canvas.fillRect(blockX*gridw/width, blockY*gridh/height, gridw/width - 2, gridh/height - 2);
                loss();
                
            }
            else if (array[blockX][blockY] == 0)
            {
                canvas.setColor(Color.white);
                canvas.fillRect(blockX*gridw/width, blockY*gridh/height, gridw/width - 2, gridh/height - 2);
                recursiveFind(blockX, blockY);
            }
        }
    }
    
    private void win()
    {
    	//recalculate
    	wins++;
    	long winTime = System.currentTimeMillis() - currentStartTime;
    	double doubleTime = (double) winTime / 1000.0;
    	if(doubleTime < bwt)
    	{
    		bwt = doubleTime;
    	}
    	awt += (doubleTime - awt) / wins;
    	win = true;
        start();
    }
    
    private void loss()
    {
    	//recalculate
    	losses++;
    	
    	win = false;
        start();
    }
    //creates the minefield
    private void createMine(int nox, int noy)
    {
        for (int n = 0; n < mines; n++)
        {
            boolean loop = true;
            int x;
            int y;
            do
            {
                x = (int) (width*Math.random());
                y = (int) (height*Math.random());
                if (array[x][y] != -1 && Math.hypot(Math.abs(x - nox), Math.abs(y - noy)) > 1.5)
                {
                    array[x][y] = -1;
                    loop = false;
                }
            }
            while (loop);
            
            int[][] adjacent = {{x - 1, y}, {x + 1, y}, {x, y + 1}, {x, y - 1}, {x - 1, y - 1}, {x - 1, y + 1}, {x + 1, y - 1}, {x + 1, y + 1}};
            
            for (int count = 0; count < 8; count++)
            {
                int dx = adjacent[count][0];
                int dy = adjacent[count][1];
                if (dx >= 0 && dx < width && dy >= 0 && dy < height && array[dx][dy] != -1)
                {
                    array[dx][dy]++;
                }
            }
        }
    }
    
    private void recursiveFind(int x, int y)
    {
        int[][] adjacent = {{x - 1, y}, {x + 1, y}, {x, y + 1}, {x, y - 1}, {x - 1, y - 1}, {x - 1, y + 1}, {x + 1, y - 1}, {x + 1, y + 1}};
        for (int count = 0; count < 8; count++)
        {
            int dx = adjacent[count][0];
            int dy = adjacent[count][1];
            if (dx >= 0 && dx < width && dy >= 0 && dy < height)
            {
                reveal(dx, dy);
            }
        }
    }
    
    private void winCheck()
    {
        int counter = 0;
        for (int w = 0; w < width; w++)
        {
            for (int h = 0; h < height; h++)
            {
                if (!revealed[w][h])
                {
                    counter++;
                }
            }
        }
        if (counter == mines)
        {
            win();
        }
    }
               
    public void mouseReleased(MouseEvent e)
    {
         if(e.getButton() == MouseEvent.BUTTON1)
         {
             reveal(e.getX()/ (gridw/width), e.getY()/ (gridh/height));
             winCheck();
         }
         
         if(e.getButton() == MouseEvent.BUTTON3 && !revealed[e.getX()/ (gridw/width)][e.getY()/ (gridh/height)])
         {
             flagged[e.getX()/ (gridw/width)][e.getY()/ (gridh/height)] = !flagged[e.getX()/ (gridw/width)][e.getY()/ (gridh/height)];
             if(flagged[e.getX()/ (gridw/width)][e.getY()/ (gridh/height)])
             {
                canvas.setColor(Color.pink);
             }
             else
             {
                 if(win)
                 {
                     canvas.setColor(Color.blue);
                 }
                 else
                 {
                     canvas.setColor(Color.red);
                 }
             }
             canvas.fillRect((e.getX()/ (gridw/width))*gridw/width, (e.getY()/ (gridh/height))*gridh/height, gridw/width - 2, gridh/height - 2);
        }
    }
    
    public void mousePressed(MouseEvent e) {}
    public void mouseClicked(MouseEvent e) {}
    public void mouseEntered(MouseEvent e) {}
    public void mouseExited(MouseEvent e) {}
    
    public static void main(String[] args)
    {
        Mine gl =  new Mine();
        gl.setup();
    }

	public void windowClosing(WindowEvent arg0)
	{
		String output = wins + " " + losses + " " + awt + " " + bwt;
		
		char[] finaloutput = output.toCharArray();
		try {
			BufferedWriter out = new BufferedWriter(new FileWriter("records.txt"));
			out.write(finaloutput);
			out.close();
			
		}
		catch (IOException e) {};
		frame.dispose();
		System.exit(0);
	}

	public void windowDeactivated(WindowEvent arg0) {

	}

	public void windowDeiconified(WindowEvent arg0) {

	}
	
	public void windowIconified(WindowEvent arg0) {

	}
	
	public void windowOpened(WindowEvent arg0) {

	}
	
	public void windowClosed(WindowEvent arg0) {
		
	}
	
	public void windowActivated(WindowEvent arg0) {

	}
}
